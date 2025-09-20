# ---
# jupyter:
#   jupytext:
#     text_representation:
#       extension: .py
#       format_name: percent
#       format_version: '1.3'
#       jupytext_version: 1.17.3
#   kernelspec:
#     display_name: Python 3 (ipykernel)
#     language: python
#     name: python3
# ---

# %%
import polars as pl

# %%
bikeshare = pl.DataFrame(dict(label=["olin", "wellesley"], num_of_bikes=[10, 2]))
bikeshare

# %%
bikeshare[0]

# %%
bikeshare.filter(bikeshare["label"] == "wellesley")

# %%
bikeshare["num_of_bikes"]


# %%
def reassign_column(df, label, column, value):
    return df.with_columns(
        pl.when(df["label"] == label).then(value).otherwise(df[column]).alias(column)
    )


# bikeshare.with_columns(pl.when(bikeshare["label"] == "olin").then(9).otherwise(bikeshare["num_of_bikes"]).alias("num_of_bikes"))
reassign_column(bikeshare, "olin", "num_of_bikes", 9)


# %%
def bike_from_x_to_y(df, *, x, y):
    df = reassign_column(df, x, "num_of_bikes", bikeshare["num_of_bikes"] - 1)
    df = reassign_column(df, y, "num_of_bikes", bikeshare["num_of_bikes"] + 1)
    return df


def bike_to_wellesley(df):
    return bike_from_x_to_y(df, x="olin", y="wellesley")


def bike_to_olin(df):
    return bike_from_x_to_y(df, x="wellesley", y="olin")


bike_to_wellesley(bikeshare)

# %%
import numpy as np


def flip(p=0.5):
    return np.random.random() < p


# %%
def step(df, p1, p2):
    if flip(p1):
        df = bike_to_wellesley(df)

    if flip(p2):
        df = bike_to_olin(df)

    return df


# %%
step(bikeshare, 0.5, 0.33)

# %%
