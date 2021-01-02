
const BaseEndPoint = "http://104.243.44.32:8080/api/v0/"

const UploadEndPoint += "upload"

const GetEndPoint += "get"

/**
 * @author Jviguy
 *
 * @param url {string} - the said url to be uploaded.
 *
 * @returns {string}
 */
async function UploadLink(url) {
    let json = await req("POST", UploadEndPoint, JSON.stringify({"url": url}))
    return json.url
}

/**
 * @author Jviguy
 *
 * @param id {string} - the said link id
 */
async function GetData(id) {
    return await req("GET", GetEndPoint, JSON.stringify({"id": id}))
}

function req(method, url, data) {
    return new Promise(function (resolve, reject) {
        let xhr = new XMLHttpRequest();
        xhr.open(method, url);
        xhr.setRequestHeader('Content-type', 'application/json');
        xhr.onload = function () {
            if (this.status >= 200 && this.status < 300) {
                resolve(xhr.response);
            } else {
                reject({
                    status: this.status,
                    statusText: xhr.statusText
                });
            }
        };
        xhr.onerror = function () {
            reject({
                status: this.status,
                statusText: xhr.statusText
            });
        };
        xhr.send();
    });
}
