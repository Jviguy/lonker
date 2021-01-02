
const BaseEndPoint = "http://104.243.44.32:8080/api/v0/"

const UploadEndPoint += "upload"

/**
 * @author Jviguy
 *
 * @param {string} - the said url to be uploaded.
 *
 * @returns {string}
 */
async function UploadLink(url) {
    return await req("POST", UploadEndPoint, JSON.stringify({"url": url}))
}

function req(method, url, data) {
    return new Promise(function (resolve, reject) {
        let xhr = new XMLHttpRequest();
        xhr.open(method, url);
        http.setRequestHeader('Content-type', 'application/json');
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
