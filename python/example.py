from university_sdk import UniversitySDK

if __name__ == "__main__":
    sdk = UniversitySDK("../data/universities.json")

    print("所有学校:")
    for u in sdk.list_all():
        print(u)

    print("\n搜索‘北京’的学校:")
    for u in sdk.search_by_province('北京'):
        print(u)

    print("\n搜索名称包含‘工业’的学校:")
    for u in sdk.search_by_name('工业'):
        print(u)
