export function timeChange(time:string) {
    var date = time.substr(0, 10); //年月日
    var hours = time.substring(11, 13);
    var minutes = time.substring(14, 16);
    var seconds = time.substring(17, 19);
    var timeFlag :any = date + ' ' + hours + ':' + minutes + ':' + seconds;
    timeFlag = timeFlag.replace(/-/g, "/");
    timeFlag = new Date(timeFlag);
    // timeFlag = new Date(timeFlag.getTime() + 8 * 3600 * 1000);
    timeFlag = new Date(timeFlag.getTime() );
    timeFlag = timeFlag.getFullYear() + '-' + ((timeFlag.getMonth() + 1) < 10 ? "0" + (timeFlag.getMonth() + 1) : (timeFlag.getMonth() + 1)) + '-' + (timeFlag.getDate() < 10 ? "0" + timeFlag.getDate() : timeFlag.getDate()) + ' ' + timeFlag.getHours() + ':' + timeFlag.getMinutes() + ':' + (timeFlag.getSeconds() < 10 ? "0" + timeFlag.getSeconds() : timeFlag.getSeconds());
    return timeFlag;
}