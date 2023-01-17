const path = require('path')
var os = require('os');
let basedir = __dirname;
const filename = path.basename(basedir);

//open ffi
const ffi = require('ffi-napi');
const ref = require("ref-napi");

//uploaderの実行フォルダー　Platform対応
function filedirPlatform(){
    const is_windows = process.platform==='win32'
    const is_mac = process.platform==='darwin'
    let libname = "";
    if(is_windows){
      libname = "libTest.dll"
      return libname
    } else if(is_mac){
      libname = "./libTest.dylib"
      return libname
    } else {
      libname = "libTest.so"
      return libname
    }
  };

const result = filedirPlatform();
console.log("libPath  > " + result)

var ffilib = ffi.Library(result ,{
    "get" : [ref.types.CString,["string"]],
  });

let mesResult = ffilib.get("hello");
console.log("calldata Message > " + mesResult);
