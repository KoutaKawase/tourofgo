from PIL import Image
import os

PERCENT = 17
DIR = "../assets/base"
OUTPUT_DIR = "../assets/cards"
MAX = 52



def main():
    for i in range(MAX):
        file_name = f"card{i+1}.png"
        p = os.path.join(DIR, file_name)
        out = os.path.join(OUTPUT_DIR, file_name)
        img = Image.open(p)
        size = (int(img.width * PERCENT / 100), int(img.height * PERCENT / 100))
        resized_img = img.resize(size)
        resized_img.save(out)


if __name__ == "__main__":
    main()
