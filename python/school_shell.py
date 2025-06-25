import requests
import json

def fetch_universities():
    url = "https://api.hcfpz.cn/un/schools"
    resp = requests.get(url)
    data = resp.json()
    with open("universities_full.json", "w", encoding="utf-8") as f:
        json.dump(data, f, ensure_ascii=False, indent=2)
    print(f"共保存 {len(data)} 所高校数据到 universities_full.json")

if __name__ == "__main__":
    fetch_universities()
