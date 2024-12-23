function x64Add(e, t) {
    var n = [e[0] >>> 16, 65535 & e[0], e[1] >>> 16, 65535 & e[1]];
    var r = [t[0] >>> 16, 65535 & t[0], t[1] >>> 16, 65535 & t[1]];
    var o = [0, 0, 0, 0];
    o[3] += n[3] + r[3];
    o[2] += o[3] >>> 16;
    o[3] &= 65535;
    o[2] += n[2] + r[2];
    o[1] += o[2] >>> 16;
    o[2] &= 65535;
    o[1] += n[1] + r[1];
    o[0] += o[1] >>> 16;
    o[1] &= 65535;
    o[0] += n[0] + r[0];
    o[0] &= 65535;
    return [o[0] << 16 | o[1], o[2] << 16 | o[3]];
}

function x64Multiply(e, t) {
    var n = [e[0] >>> 16, 65535 & e[0], e[1] >>> 16, 65535 & e[1]];
    var r = [t[0] >>> 16, 65535 & t[0], t[1] >>> 16, 65535 & t[1]];
    var o = [0, 0, 0, 0];
    o[3] += n[3] * r[3];
    o[2] += o[3] >>> 16;
    o[3] &= 65535;
    o[2] += n[2] * r[3];
    o[1] += o[2] >>> 16;
    o[2] &= 65535;
    o[2] += n[3] * r[2];
    o[1] += o[2] >>> 16;
    o[2] &= 65535;
    o[1] += n[1] * r[3];
    o[0] += o[1] >>> 16;
    o[1] &= 65535;
    o[1] += n[2] * r[2];
    o[0] += o[1] >>> 16;
    o[1] &= 65535;
    o[1] += n[3] * r[1];
    o[0] += o[1] >>> 16;
    o[1] &= 65535;
    o[0] += n[0] * r[3] + n[1] * r[2] + n[2] * r[1] + n[3] * r[0];
    o[0] &= 65535;
    return [o[0] << 16 | o[1], o[2] << 16 | o[3]];
}

function x64Rotl(e, t) {
    var n = t % 64;
    return 32 === n ? [e[1], e[0]] : n < 32 ? [e[0] << n | e[1] >>> 32 - n, e[1] << n | e[0] >>> 32 - n] : (n -= 32, [e[1] << n | e[0] >>> 32 - n, e[0] << n | e[1] >>> 32 - n]);
}

function x64LeftShift(e, t) {
    var n = t % 64;
    return 0 === n ? e : n < 32 ? [e[0] << n | e[1] >>> 32 - n, e[1] << n] : [e[1] << n - 32, 0];
}

function x64Xor(e, t) {
    return [e[0] ^ t[0], e[1] ^ t[1]];
}

function x64Fmix(t) {
    var n = t;
    n = x64Xor(n, [0, n[0] >>> 1]);
    n = x64Multiply(n, [4283543511, 3981806797]);
    n = x64Xor(n, [0, n[0] >>> 1]);
    n = x64Multiply(n, [3301882366, 444984403]);
    return n = x64Xor(n, [0, n[0] >>> 1]);
}

function x64hash128(t, n) {
    var r = t || "";
    var o = n || 0;
    var i = r.length % 16;
    var a = r.length - i;
    var c = [0, o];
    var u = [0, o];
    var l = [0, 0];
    var s = [0, 0];
    var f = [2277735313, 289559509];
    var h = [1291169091, 658871167];
    for (var d = 0; d < a; d += 16) {
        l = [255 & r.charCodeAt(d + 4) | (255 & r.charCodeAt(d + 5)) << 8 | (255 & r.charCodeAt(d + 6)) << 16 | (255 & r.charCodeAt(d + 7)) << 24, 255 & r.charCodeAt(d) | (255 & r.charCodeAt(d + 1)) << 8 | (255 & r.charCodeAt(d + 2)) << 16 | (255 & r.charCodeAt(d + 3)) << 24];
        s = [255 & r.charCodeAt(d + 12) | (255 & r.charCodeAt(d + 13)) << 8 | (255 & r.charCodeAt(d + 14)) << 16 | (255 & r.charCodeAt(d + 15)) << 24, 255 & r.charCodeAt(d + 8) | (255 & r.charCodeAt(d + 9)) << 8 | (255 & r.charCodeAt(d + 10)) << 16 | (255 & r.charCodeAt(d + 11)) << 24];
        l = x64Multiply(l, f);
        l = x64Rotl(l, 31);
        l = x64Multiply(l, h);
        c = x64Xor(c, l);
        c = x64Rotl(c, 27);
        c = x64Add(c, u);
        c = x64Add(x64Multiply(c, [0, 5]), [0, 1390208809]);
        s = x64Multiply(s, h);
        s = x64Rotl(s, 33);
        s = x64Multiply(s, f);
        u = x64Xor(u, s);
        u = x64Rotl(u, 31);
        u = x64Add(u, c);
        u = x64Add(x64Multiply(u, [0, 5]), [0, 944331445]);
    }
    l = [0, 0];
    s = [0, 0];
    switch (i) {
        case 15:
            s = x64Xor(s, x64LeftShift([0, r.charCodeAt(d + 14)], 48));
            break;
        case 14:
            s = x64Xor(s, x64LeftShift([0, r.charCodeAt(d + 13)], 40));
            break;
        case 13:
            s = x64Xor(s, x64LeftShift([0, r.charCodeAt(d + 12)], 32));
            break;
        case 12:
            s = x64Xor(s, x64LeftShift([0, r.charCodeAt(d + 11)], 24));
            break;
        case 11:
            s = x64Xor(s, x64LeftShift([0, r.charCodeAt(d + 10)], 16));
            break;
        case 10:
            s = x64Xor(s, x64LeftShift([0, r.charCodeAt(d + 9)], 8));
            break;
        case 9:
            s = x64Xor(s, [0, r.charCodeAt(d + 8)]);
            s = x64Multiply(s, h);
            s = x64Rotl(s, 33);
            s = x64Multiply(s, f);
            u = x64Xor(u, s);
            break;
        case 8:
            l = x64Xor(l, x64LeftShift([0, r.charCodeAt(d + 7)], 56));
            break;
        case 7:
            l = x64Xor(l, x64LeftShift([0, r.charCodeAt(d + 6)], 48));
            break;
        case 6:
            l = x64Xor(l, x64LeftShift([0, r.charCodeAt(d + 5)], 40));
            break;
        case 5:
            l = x64Xor(l, x64LeftShift([0, r.charCodeAt(d + 4)], 32));
            break;
        case 4:
            l = x64Xor(l, x64LeftShift([0, r.charCodeAt(d + 3)], 24));
            break;
        case 3:
            l = x64Xor(l, x64LeftShift([0, r.charCodeAt(d + 2)], 16));
            break;
        case 2:
            l = x64Xor(l, x64LeftShift([0, r.charCodeAt(d + 1)], 8));
            break;
        case 1:
            l = x64Xor(l, [0, r.charCodeAt(d)]);
            l = x64Multiply(l, f);
            l = x64Rotl(l, 31);
            l = x64Multiply(l, h);
            c = x64Xor(c, l);
    }
    c = x64Xor(c, [0, r.length]);
    u = x64Xor(u, [0, r.length]);
    c = x64Add(c, u);
    u = x64Add(u, c);
    c = x64Fmix(c);
    u = x64Fmix(u);
    c = x64Add(c, u);
    u = x64Add(u, c);
    return ("00000000" + (c[0] >>> 0).toString(16)).slice(-8) + ("00000000" + (c[1] >>> 0).toString(16)).slice(-8) + ("00000000" + (u[0] >>> 0).toString(16)).slice(-8) + ("00000000" + (u[1] >>> 0).toString(16)).slice(-8);
}

function validateAgainstMask(e, t) {
    t = t || "";
    return 0 === e.substring(0, t.length).localeCompare(t);
}

function randomStringGenerator(e, t) {
    var n = "";
    for (var o = 0; o < t - e; o++) {
        n += "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789".charAt(Math.floor(Math.random() * "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789".length));
    }
    return n;
}

function solvePow(e, t, n) {
    var r = false;
    var a = "";
    do {
        a = e + randomStringGenerator(e.length, 16);
        var c = x64hash128(a, t);
        r = validateAgainstMask(c, n);
    } while (!r);
    return a;
}


var mdata = process.argv[2]
let a = JSON.parse(`${mdata}`.replace(/\\/g, ""));
console.log(solvePow(a.key, a.seed, a.mask))