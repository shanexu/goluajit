package goluajit

// #include <stdio.h>
// #include <lua.h>
// #include <lualib.h>
// #include <lauxlib.h>
import "C"

type LuaState struct {
	ls *C.lua_State
}

func NewState() *LuaState {
	ls := C.luaL_newstate()
	s := &LuaState{
		ls: ls,
	}
	return s
}

func (s *LuaState) OpenLibs() {
	C.luaL_openlibs(s.ls)
}

func (s *LuaState) LoadFile(file string) {
	C.luaL_loadfile(s.ls, C.CString(file))
}

func (s *LuaState) PCall(i1, i2, i3 int) int {
	return int(C.lua_pcall(s.ls, C.int(i1), C.int(i2), C.int(i3)))
}

func (s *LuaState) GetGlobal(name string) {
	C.lua_getfield(s.ls, C.LUA_GLOBALSINDEX, C.CString(name))
}

func (s *LuaState) ToString(i int) string {
	return C.GoString(C.lua_tolstring(s.ls, C.int(i), nil))
}

func (s *LuaState) ToInteger(i int) int {
	return int(C.lua_tointeger(s.ls, C.int(i)))
}

func (s *LuaState) SetTop(i int) {
	C.lua_settop(s.ls, C.int(i))
}
