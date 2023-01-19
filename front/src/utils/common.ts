export function timeChange(time: string) {
    var date = time.substr(0, 10); //年月日
    var hours = time.substring(11, 13);
    var minutes = time.substring(14, 16);
    var seconds = time.substring(17, 19);
    var timeFlag: any = date + ' ' + hours + ':' + minutes + ':' + seconds;
    timeFlag = timeFlag.replace(/-/g, "/");
    timeFlag = new Date(timeFlag);
    // timeFlag = new Date(timeFlag.getTime() + 8 * 3600 * 1000);
    timeFlag = new Date(timeFlag.getTime());
    timeFlag = timeFlag.getFullYear() + '-' + ((timeFlag.getMonth() + 1) < 10 ? "0" + (timeFlag.getMonth() + 1) : (timeFlag.getMonth() + 1)) + '-' + (timeFlag.getDate() < 10 ? "0" + timeFlag.getDate() : timeFlag.getDate()) + ' ' + timeFlag.getHours() + ':' + (timeFlag.getMinutes() < 10 ? "0" + timeFlag.getMinutes() : timeFlag.getMinutes()) + ':' + (timeFlag.getSeconds() < 10 ? "0" + timeFlag.getSeconds() : timeFlag.getSeconds());
    return timeFlag;
}

// 下载文件，并自动更换文件名
export const downLoadFile = (url: string, name: string) => {

    let getBlob = function (url: string, cb: any) { // 获取文件流
        let xhr = new XMLHttpRequest()
        xhr.open('GET', url, true)
        xhr.responseType = 'blob'
        xhr.onload = function () {
            if (xhr.status === 200) {
                cb(xhr.response)
            }
        }
        xhr.send()
    }

    let saveAs = function (blob: any, filename: string) { // 改名字
        let link = document.createElement('a')
        let body = document.querySelector('body')

        link.href = window.URL.createObjectURL(blob)
        link.download = filename

        // fix Firefox
        link.style.display = 'none'
        body!.appendChild(link)

        link.click()
        body!.removeChild(link)

        window.URL.revokeObjectURL(link.href)
    }
    let download = function (url: string, filename: string) { // 执行
        getBlob(url, function (blob: any) {
            saveAs(blob, filename)
        })
    };

    // 下载文件
    download(url, name) // OBS可下载的文件url，你想要改的名字

}


export const OBS_URL: string = "https://file-report-store.obs.cn-north-4.myhuaweicloud.com/"
