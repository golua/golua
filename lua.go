package golua

//struct lua_Debug {
//int event;
//const char *name;	/* (n) */
//const char *namewhat;	/* (n) 'global', 'local', 'field', 'method' */
//const char *what;	/* (S) 'Lua', 'C', 'main', 'tail' */
//const char *source;	/* (S) */
//int currentline;	/* (l) */
//int linedefined;	/* (S) */
//int lastlinedefined;	/* (S) */
//unsigned char nups;	/* (u) number of upvalues */
//unsigned char nparams;/* (u) number of parameters */
//char isvararg;        /* (u) */
//char istailcall;	/* (t) */
//char short_src[LUA_IDSIZE]; /* (S) */
///* private part */
//struct CallInfo *i_ci;  /* active function */
//};

/*
** basic types
 */
const (
	LUA_TNONE = -1

	LUA_TNIL           = 0
	LUA_TBOOLEAN       = 1
	LUA_TLIGHTUSERDATA = 2
	LUA_TNUMBER        = 3
	LUA_TSTRING        = 4
	LUA_TTABLE         = 5
	LUA_TFUNCTION      = 6
	LUA_TUSERDATA      = 7
	LUA_TTHREAD        = 8

	LUA_NUMTAGS = 9
)

/* thread status */
const (
	LUA_OK        = 0
	LUA_YIELD     = 1
	LUA_ERRRUN    = 2
	LUA_ERRSYNTAX = 3
	LUA_ERRMEM    = 4
	LUA_ERRGCMM   = 5
	LUA_ERRERR    = 6
)

/*
** Comparison and arithmetic functions
 */
const (
	LUA_OPADD  = 0
	LUA_OPSUB  = 1
	LUA_OPMUL  = 2
	LUA_OPMOD  = 3
	LUA_OPPOW  = 4
	LUA_OPDIV  = 5
	LUA_OPIDIV = 6
	LUA_OPBAND = 7
	LUA_OPBOR  = 8
	LUA_OPBXOR = 9
	LUA_OPSHL  = 10
	LUA_OPSHR  = 11
	LUA_OPUNM  = 12
	LUA_OPBNOT = 13
)

const (
	LUA_OPEQ = 0
	LUA_OPLT = 1
	LUA_OPLE = 2
)

type LuaDebug struct {
	Event       int32
	Name        string
	Namewhat    string
	What        string
	Source      string
	Currentline int32
	Linedefined int32
	Nups        int8
	Nparams     int8
	Isvararg    bool
	Istailcall  bool
	Shortsrce   []byte
}

/*
** basic stack manipulation
*/
//LUA_API int   (lua_absindex) (lua_State *L, int idx);
//LUA_API int   (lua_gettop) (lua_State *L);
//LUA_API void  (lua_settop) (lua_State *L, int idx);
//LUA_API void  (lua_pushvalue) (lua_State *L, int idx);
//LUA_API void  (lua_rotate) (lua_State *L, int idx, int n);
//LUA_API void  (lua_copy) (lua_State *L, int fromidx, int toidx);
//LUA_API int   (lua_checkstack) (lua_State *L, int n);
//
//LUA_API void  (lua_xmove) (lua_State *from, lua_State *to, int n);

func Lua_absindex(L *interface{},idx int) int {
	//TODO
	return 	idx
}

func main() {
}
