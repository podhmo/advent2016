# -*- coding:utf-8 -*-
# see: https://mholt.github.io/json-to-go/
# original: https://github.com/mholt/json-to-go/blob/master/json-to-go.js
import re
import json
from collections import defaultdict
from collections import deque
from prestring import NameStore
from prestring import PreString
from prestring import LazyFormat
from prestring.go import GoModule
from prestring.go import goname as to_goname


def json_to_go(json_string, name, m=None, rx=re.compile("\.0", re.M)):
    m = m or GoModule()
    data = json.loads(rx.sub(".1", json_string))
    s = detect_struct_info(data, name)

    with m.import_group() as im:
        pass
    emit_code(s, name, m=m, im=im)
    im.clear_ifempty()
    return m


def resolve_type(val, time_rx=re.compile("\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d(\.\d+)?(\+\d\d:\d\d|Z)")):
    if val is None:
        return "interface{}"
    if isinstance(val, str):
        if time_rx.match(val):
            return "time.Time"
        elif "://" in val:
            return "github.com/go-openapi/strfmt.Uri"
        else:
            return "string"
    elif isinstance(val, int):
        if val > -2147483648 and val < 2147483647:
            return "int"
        else:
            return "int64"
    elif isinstance(val, float):
        return "float64"
    elif isinstance(val, bool):
        return "bool"
    elif hasattr(val, "keys"):
        return "struct"
    elif isinstance(val, (list, tuple)):
        return "slice"
    else:
        raise ValueError("unsupported for {!r}".format(val))


def select_better_type(*types):
    s = {t for t in types if t is not None}
    if "float64" in s:
        return "float64"
    elif "int64" in s:
        return "int64"
    else:
        return s.pop()


def detect_struct_info(d, name):
    def _detect_struct_info(d, s, name):
        if hasattr(d, "keys"):
            s["type"] = "struct"
            s["jsonname"] = name
            s["freq"] += 1
            for k, v in d.items():
                goname = to_goname(k)
                _detect_struct_info(v, s["children"][goname], k)
        elif isinstance(d, (list, tuple)):
            s["type2"] = "slice"
            for x in d:
                _detect_struct_info(x, s, name)  # xxx
        else:
            typ = resolve_type(d)
            s["jsonname"] = name
            s["freq"] += 1
            s["type"] = select_better_type(s["type"], typ)

    def make_struct_info():
        return {"freq": 0, "type": None, "children": defaultdict(make_struct_info)}
    s = defaultdict(make_struct_info)
    goname = to_goname(name)
    _detect_struct_info(d, s[goname], goname)
    return s[goname]


def to_type_struct_info(sinfo):
    if sinfo.get("type2") == "slice":
        return "[]" + sinfo["type"]
    else:
        return sinfo["type"]


def is_omitempty_struct_info(subinfo, sinfo):
    return subinfo["freq"] < sinfo["freq"]


def emit_code(sinfo, name, m, im, name_score_map={"parent": -1, '': -10}):
    cw = CommentWriter(m, name, sinfo)
    ns = NameStore()
    defs = set()
    typename_map = defaultdict(lambda: PreString(""))

    def make_signature(sinfo):
        return tuple([(k, v["type"], v.get("type2")) for k, v in sorted(sinfo["children"].items())])

    def emit_structure_comment(sinfo, name, parent=None):
        if sinfo.get("type") == "struct":
            cw.write(name, sinfo, parent=parent)
            for name, subinfo in sorted(sinfo["children"].items()):
                emit_structure_comment(subinfo, name, parent=sinfo)

    def _emit_struct(sinfo, name, parent=None):
        with m.type_(name, to_type_struct_info(sinfo)):
            for name, subinfo in sorted(sinfo["children"].items()):
                _emit_code(subinfo, name, m, parent=sinfo)

    def _emit_code(sinfo, name, m, parent=None):
        if "." in sinfo.get("type"):
            im.import_(sinfo.get("type").rsplit(".", 1)[0])

        if sinfo.get("type") == "struct":
            signature = make_signature(sinfo)
            cont.append((name, sinfo, signature))
            typ = typename_map[signature]
            typ.body.append(name)
        else:
            typ = to_type_struct_info(sinfo)
            if "/" in typ:
                typ = typ.rsplit("/", 1)[-1]
        m.stmt(LazyFormat('{} {}', name, typ))

        # append tag
        if is_omitempty_struct_info(sinfo, parent):
            m.insert_after('  `json:"{},omitempty"`'.format(sinfo["jsonname"]))
        else:
            m.insert_after('  `json:"{}"`'.format(sinfo["jsonname"]))

    emit_structure_comment(sinfo, name, parent=None)

    cont = deque([(name, sinfo, make_signature(sinfo))])
    while cont:
        name, sinfo, signature = cont.popleft()
        if signature in defs:
            continue
        defs.add(signature)
        typename_map[signature].body.append(name)
        _emit_struct(sinfo, typename_map[signature])

    for signature, lazy_typename in typename_map.items():
        candidates = set(lazy_typename.body)
        new_name = max(candidates, key=lambda k: name_score_map.get(k.lower(), 0))
        ns[signature] = new_name
        lazy_typename.body.clear()
        lazy_typename.body.append(ns[signature])
    return m


class CommentWriter(object):
    def __init__(self, m, name, sinfo):
        m.stmt("/* structure")
        cm = GoModule()
        m.stmt(cm)
        cm.stmt(name)
        self.cm_map = {sinfo["jsonname"]: cm}
        m.stmt("*/")

    def write(self, name, sinfo, parent=None):
        if parent is None:
            return
        cm = self.cm_map[parent["jsonname"]]
        with cm.scope():
            cm.stmt(name)
            self.cm_map[sinfo["jsonname"]] = cm.submodule(newline=False)


if __name__ == "__main__":
    import argparse
    parser = argparse.ArgumentParser()
    parser.add_argument("--package", type=str, default="autogen")
    parser.add_argument("--name", type=str, default="AutoGenerated")
    parser.add_argument("src", type=argparse.FileType('r'))
    args = parser.parse_args()

    m = GoModule()
    m.package(args.package)
    print(json_to_go(args.src.read(), args.name, m))
