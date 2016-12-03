from prestring.go import Module

m = Module()
m.package('main')
with m.import_group() as im:
    pass

with m.func('main'):
    im.import_('log')
    m.stmt('log.Println("hmm")')

print(m)
