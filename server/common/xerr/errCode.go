package xerr

/* 全局错误码 */
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/
//全局错误码
const SERVER_COMMON_ERROR uint32 = 100001
const REUQEST_PARAM_ERROR uint32 = 100002

const DB_ERROR uint32 = 100005
const DB_UPDATE_AFFECTED_ZERO_ERROR uint32 = 100006

const DB_NOTFOUNR = 100100
const DB_DELETE_NOTEXIST = 100101 // 删除数据时数据库不存在此条数据
const FILE_UPLOAD_ERROR = 100300

const TOKEN_NOT_EXIST uint32 = 200001
const TOKEN_EXPIRE_ERROR uint32 = 200003
const TOKEN_GENERATE_ERROR uint32 = 200004
