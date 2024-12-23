import base64
import json
import math
import random
import subprocess

import tls_client
from bs4 import BeautifulSoup

CPID = "6e019ad0-8486-4923-a5bd-35f6363d75fd"


def generate() -> dict:
    # TODO 実在するデータ使ってランダムに生成させる
    return {
        "fonts": {
            "value": [
                "sans-serif-thin",
                "ARNO PRO",
                "Agency FB",
                "Arabic Typesetting",
                "Arial Unicode MS",
                "AvantGarde Bk BT",
                "BankGothic Md BT",
                "Batang",
                "Bitstream Vera Sans Mono",
                "Calibri",
                "Century",
                "Century Gothic",
                "Clarendon",
                "EUROSTILE",
                "Franklin Gothic",
                "Futura Bk BT",
                "Futura Md BT",
                "GOTHAM",
                "Gill Sans",
                "HELV",
                "Haettenschweiler",
                "Helvetica Neue",
                "Humanst521 BT",
                "Leelawadee",
                "Letter Gothic",
                "Levenim MT",
                "Lucida Bright",
                "Lucida Sans",
                "Menlo",
                "MS Mincho",
                "MS Outlook",
                "MS Reference Specialty",
                "MS UI Gothic",
                "MT Extra",
                "MYRIAD PRO",
                "Marlett",
                "Meiryo UI",
                "Microsoft Uighur",
                "Minion Pro",
                "Monotype Corsiva",
                "PMingLiU",
                "Pristina",
                "SCRIPTINA",
                "Segoe UI Light",
                "Serifa",
                "SimHei",
                "Small Fonts",
                "Staccato222 BT",
                "TRAJAN PRO",
                "Univers CE 55 Medium",
                "Vrinda",
                "ZWAdobeF",
            ],
            "duration": random.randint(150, 170),
        },
        "domBlockers": {"duration": 16},
        "fontPreferences": {
            "value": {
                "default": 164.53326021245005,
                "apple": 164.53326021245005,
                "serif": 164.53326021245005,
                "sans": 155.76382478978442,
                "mono": 122.40449322825657,
                "min": 10.329785612962995,
                "system": 164.53326021245005,
            },
            "duration": random.randint(150, 220),
        },
        "audio": {"value": 35.74996319835191, "duration": random.randint(0, 2)},
        "screenFrame": {"value": [0, 0, 0, 0], "duration": 0},
        "osCpu": {"value": "Windows NT 6.2; Win64; x64", "duration": 0},
        "languages": {"value": [["en-US"], ["en-US", "en"]], "duration": 0},
        "colorDepth": {"value": 24, "duration": 0},
        "deviceMemory": {"duration": 0},
        "screenResolution": {"value": [1920, 1080], "duration": 0},
        "hardwareConcurrency": {"value": 4, "duration": 0},
        "timezone": {"value": "WET", "duration": 0},
        "sessionStorage": {"value": True, "duration": 0},
        "localStorage": {"value": True, "duration": 0},
        "indexedDB": {"value": True, "duration": 0},
        "openDatabase": {"value": False, "duration": 0},
        "cpuClass": {"duration": 0},
        "platform": {"value": "Win32", "duration": 0},
        "plugins": {"value": [], "duration": 0},
        "canvas": {
            "value": {"winding": True, "geometry": "unstable", "text": "unstable"},
            "duration": 0,
        },
        "touchSupport": {
            "value": {"maxTouchPoints": 0, "touchEvent": False, "touchStart": False},
            "duration": 0,
        },
        "vendor": {"value": "", "duration": 0},
        "vendorFlavors": {"value": [], "duration": 0},
        "cookiesEnabled": {"value": True, "duration": 0},
        "colorGamut": {"value": "srgb", "duration": 0},
        "invertedColors": {"duration": 0},
        "forcedColors": {"value": False, "duration": 0},
        "monochrome": {"value": 0, "duration": 0},
        "contrast": {"value": 0, "duration": 0},
        "reducedMotion": {"value": False, "duration": 0},
        "hdr": {"value": False, "duration": 0},
        "math": {
            "value": {
                "acos": math.acos(0.12312423423423424),
                "acosh": math.acosh(1e308),
                "acoshPf": math.log(1e154 + math.sqrt(1e308)),
                "asin": math.asin(0.12312423423423424),
                "asinh": math.asinh(1),
                "asinhPf": math.log(1 + math.sqrt(1 * 1 + 1)),
                "atanh": math.atanh(0.5),
                "atanhPf": math.log((1 + 0.5) / (1 - 0.5)) / 2,
                "atan": math.atan(0.5),
                "sin": math.sin(-1e300),
                "sinh": math.sinh(1),
                "sinhPf": math.exp(1) - 1 / math.exp(1) / 2,
                "cos": math.cos(10.000000000123),
                "cosh": math.cosh(1),
                "coshPf": (math.exp(1) + 1 / math.exp(1)) / 2,
                "tan": math.tan(-1e300),
                "tanh": math.tanh(1),
                "tanhPf": (math.exp(2 * 1) - 1) / (math.exp(2 * 1) + 1),
                "exp": math.exp(1),
                "expm1": math.expm1(1),
                "expm1Pf": math.exp(1) - 1,
                "log1p": math.log1p(1),
                "log1pPf": math.log(1 + 10),
                "powPI": math.pow(math.pi, -100),
            },
            "duration": 0,
        },
        "videoCard": {"duration": random.randint(2, 150)},
        "pdfViewerEnabled": {"value": True, "duration": 0},
        "architecture": {"value": 255, "duration": 0},
    }


def ask_send(
    shop_id: str, email_address: str, message: str, last_name: str, first_name: str
):
    session = tls_client.Session(
        client_identifier="firefox_120", random_tls_extension_order=True
    )

    components = generate()
    result = subprocess.run(
        ["node", "components_hash.js", f"{json.dumps(components)}"],
        capture_output=True,
        text=True,
    )
    components_hash = result.stdout.strip()
    headers = {
        "Content-Type": "text/plain; charset=utf-8",
        "Origin": "https://ask.step.rakuten.co.jp",
        "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0",
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
        "Accept-Language": "en-US,en;q=0.5",
        "Sec-GPC": "1",
        "Connection": "keep-alive",
        "Upgrade-Insecure-Requests": "1",
        "Sec-Fetch-Dest": "document",
        "Sec-Fetch-Mode": "navigate",
        "Sec-Fetch-Site": "none",
        "Sec-Fetch-User": "?1",
        "Priority": "u=0, i",
        "Pragma": "no-cache",
        "Cache-Control": "no-cache",
    }

    response = session.post(
        "https://challenger.api.global.rakuten.com/v1.0/c",
        headers=headers,
        json={
            "pid": CPID,
            "lang": "ja-JP",
            "param": {},
            "rat": {"hash": components_hash, "components": components.pop("canvas")},
        },
    )
    print(response)
    cid = response.json()["result"]["cid"]

    response = session.get(
        "https://challenger.api.global.rakuten.com/v1.0/m",
        params={"cid": cid, "mtype": "0"},
        headers=headers,
    )
    print(response)
    mdata = response.json()["media"][0]["data"]

    result = subprocess.run(
        ["node", "cres.js", f"{mdata}"], capture_output=True, text=True
    )
    pow = result.stdout.strip()

    data = json.loads(
        str(response.json()["media"][1]["data"])
        .replace("/\\/g", "")
        .replace("\\", " ")
        .replace(" ", "")
    )
    for key in data.keys():
        if data[key]["input"] == "checkbox" and data[key]["render"] == 0:
            checkbox_key = key
    print(pow, checkbox_key)

    cres = base64.b64encode(json.dumps({"cres": [pow, checkbox_key]}).encode()).decode()
    print(cres)
    headers["Referer"] = (
        f"https://ask.step.rakuten.co.jp/inquiry-form/?page=simple-inquiry-top&act=login&ms=500&shop_id={shop_id}"
    )

    response = session.get(
        "https://ask.step.rakuten.co.jp/inquiry-form/",
        params={
            "page": "simple-inquiry-nmr",
            "act": "input",
            "ms": "500",
            "shop_id": shop_id,
            "language_type": "0",
        },
        headers=headers,
    )
    soup = BeautifulSoup(response.text, "html.parser")
    csrf_key = soup.find("input", {"name": "csrf_key"})["value"]

    headers["Content-Type"] = "application/x-www-form-urlencoded"
    response = session.post(
        "https://ask.step.rakuten.co.jp/inquiry-form/",
        headers=headers,
        data={
            "language_type": "0",
            "page": "simple-inquiry-nmr",
            "act": "send",
            "csrf_key": csrf_key,
            "shop_id": shop_id,
            "ms": "500",
            "last_name": last_name,
            "first_name": first_name,
            "mail_addr": email_address,
            "mail_addr_sub": email_address,
            "inquiry_message": message,
            "cres": cres,
            "cpid": CPID,
            "cid": cid,
        },
    )
    if "お問い合わせありがとうございました。" in response.text:
        print("お問い合わせを送信しました。")
    print(response.status_code)


if __name__ == "__main__":
    ask_send(
        shop_id="",
        email_address="",
        message="",
        last_name="",
        first_name="",
    )
