const path = require('path')
var os = require('os');
let basedir = __dirname;
const filename = path.basename(basedir);

//open ffi
const ffi = require("ffi-napi");
const ref = require("ref-napi");

//uploaderの実行フォルダー　Platform対応
function filedirPlatform(){
    const is_windows = process.platform==='win32'
    const is_mac = process.platform==='darwin'
    let libname = "";
    if(is_windows){
      libname = process.cwd() + "/" + "libTest.dll"
      return libname
    } else if(is_mac){
      res = GetAppRoot() 
      libname = res[0] + "libTest.dylib"
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

function GetAppRoot(){
  let nextPath = __filename;
  cutname = GetCurrentName(nextPath)
  let logpath = "First " + nextPath + " / " + cutname + os.EOL;
  let appsPathArray = SearchPath("Contents",__filename,logpath)
  const verPath = appsPathArray[0]
  const exePath = path.normalize(verPath + '/../../')
  const logPath = appsPathArray[1] + verPath + os.EOL + exePath + os.EOL + filename

  //DebugLog
  //const savePath = exePath + "log.txt"
  //logSave(savePath,logPath)
  return [verPath,exePath]
}

//対象の文字列のディレクトリ名が出るまで検索ß
function SearchPath(searchStr,nextPath,log){
  nextPath = path.normalize(nextPath + '/../');
  cutname = GetCurrentName(nextPath)
  if(cutname == ""){
    return [nextPath,log]
  }
  log = log + nextPath + " / " + cutname + os.EOL;
  if(cutname != searchStr){
      const array = SearchPath(searchStr,nextPath,log)
      nextPath = array[0]
      log = array[1]
  }
  return [nextPath,log]
}

//親の名前を所得するファイルならファイル名、ディレクトリならカレントディレクトr名
function GetCurrentName(str){
  let index = str.lastIndexOf(path.sep)
  if(index <= 3){
    return ""
  }

  let cutname = str.slice(index+1)
  cutname = lastDeleteSep(cutname)
  if(cutname == ""){
    index = index -1
    index = str.lastIndexOf(path.sep,index)
    cutname = str.slice(index+1)
    cutname = lastDeleteSep(cutname)
  }

  return cutname
}

//文字列の末尾からセパレータを削除する
function lastDeleteSep(str){
  //1文字取り出してチェック
  let check = str.slice(-1)
  if(check == path.sep){
    //ラスト一文字削る
    str = str.slice(0,-1)
  }
  return str
}