const CryptoJS = require('crypto-js')
const Utils = {
    //https://stackoverflow.com/questions/33914764/how-to-read-a-binary-file-with-filereader-in-order-to-hash-it-with-sha-256-in-cr
    stringHash(arrayBuffer) {
        var i8a = new Uint8Array(arrayBuffer);
        var a = [];
        for (var i = 0; i < i8a.length; i += 4) {
            a.push(i8a[i] << 24 | i8a[i + 1] << 16 | i8a[i + 2] << 8 | i8a[i + 3]);
        }
        return '0x' + CryptoJS.SHA256(CryptoJS.lib.WordArray.create(a, i8a.length)).toString();
    }
}
module.exports = Utils
