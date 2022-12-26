//open ffi
const ffi = require("ffi-napi");
const ref = require("ref-napi");

//uploaderの実行フォルダー　Platform対応
function filedirPlatform(){
    const is_windows = process.platform==='win32'
    const is_mac = process.platform==='darwin'
    const pathDat = __filename;
    if(is_windows){
      libname = process.cwd() + "/" + "libTest.dll"
      return libname
    } else if(is_mac){
      libname = GetAppRoot() + "libTest.dylib"
      return libname
    } else {
      libname = process.cwd() + "/" + "libTest.so"
      return libname
    }
  };

const result = filedirPlatform();
console.log("libPath  > " + result)

var ffilib = ffi.Library(result,{
    "ffiCheck" : [ref.types.CString,["string"]],
  });

let mesResult = ffilib.ffiCheck("hello");
console.log("calldata Message > " + mesResult);
