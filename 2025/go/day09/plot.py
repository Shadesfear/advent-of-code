# /// script
# requires-python = ">=3.12"
# dependencies = ["pandas", "matplotlib"]
# ///

import pandas as pd
import matplotlib.pyplot as plt

def main() -> None:
    df = pd.read_csv("../../inputs/day09.txt", header=None)
    plt.plot(df[0], df[1])
    plt.savefig("connected.png")

    print(df.head())
    print("Hello from plot.py!")


if __name__ == "__main__":
    main()
