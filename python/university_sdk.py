import json
from typing import List

class UniversitySDK:
    def __init__(self, data_path: str):
        with open(data_path, encoding='utf-8') as f:
            self.data = json.load(f)

    def list_all(self) -> List[dict]:
        return self.data

    def search_by_name(self, keyword: str) -> List[dict]:
        keyword_lower = keyword.lower()
        return [u for u in self.data if keyword_lower in u['name'].lower()]

    def search_by_province(self, province: str) -> List[dict]:
        province_lower = province.lower()
        return [u for u in self.data if province_lower == u['province'].lower()]

    def search_by_city(self, city: str) -> List[dict]:
        city_lower = city.lower()
        return [u for u in self.data if city_lower == u['city'].lower()]