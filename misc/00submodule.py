from prestring.go import Module


m = Module()

m.stmt("begin foo")
sm = m.submodule()
m.stmt("end foo")

m.stmt("bar")

sm.sep()
sm.stmt("-- yay!")
sm.sep()

print(m)
