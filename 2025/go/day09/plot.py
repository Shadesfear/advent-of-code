# /// script
# requires-python = ">=3.12"
# dependencies = ["pandas", "matplotlib"]
# ///

import pandas as pd
import matplotlib.pyplot as plt

def main() -> None:
    df = pd.read_csv("../../inputs/day09.txt", header=None)
    plt.plot(df[0], df[1])

    p1 = (5953, 67629)
    p2 = (94872, 50262)
    rect_x = [p1[0], p2[0], p2[0], p1[0], p1[0]]
    rect_y = [p1[1], p1[1], p2[1], p2[1], p1[1]]
    plt.plot(rect_x, rect_y, 'r-', linewidth=2)

    plt.savefig("connected.png")

    print(df.head())
    print("Hello from plot.py!")


if __name__ == "__main__":
    main()
