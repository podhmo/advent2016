from prestring import PreString
from prestring.go import GoModule

m = GoModule()
p = PreString("")
m.stmt("foo")
m.stmt(p)
m.stmt("bar")

p.body.append("*inner*")

print(m)
