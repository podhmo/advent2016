import sys  # NOQA
import argparse
import json
from collections import OrderedDict
from prestring.go import GoModule  # NOQA
from goconvert import Reader, Writer
from goconvert import builders


def run(src_file, dst_file, override_file):
    reader = Reader()
    writer = Writer()
    for name, f in [("src", src_file), ("dst", dst_file), ("override", override_file)]:
        with open(f) as rf:
            reader.read_world(json.load(rf, object_pairs_hook=OrderedDict))

    convert_module = reader.universe.create_module("convert", "github.com/podhmo/advent2016/dst/convert")
    b = builders.ConvertBuilder(reader.universe, convert_module)
    b.register_from_module(convert_module)

    src = reader.universe.find_module("github.com/podhmo/advent2016/src/github")
    dst = reader.universe.find_module("github.com/podhmo/advent2016/dst/swagger/gen/def/")
    func = b.build_struct_convert(src["AuthenticatedUser"], dst["User"])
    print(func.parent.dump(writer))


def main():
    from goconvert import logger as l
    parser = argparse.ArgumentParser()
    parser.add_argument("--src", required=True)
    parser.add_argument("--dst", required=True)
    parser.add_argument("--override", required=True)
    parser.add_argument("--logger", required=False, choices=l.LEVELS, default=None)
    args = parser.parse_args()
    if args.logger:
        l.activate_logger(level=args.logger)
    return run(args.src, args.dst, args.override)

if __name__ == "__main__":
    main()
    # run("../json/src.json", "../json/dst.json", "../json/convert.json")
