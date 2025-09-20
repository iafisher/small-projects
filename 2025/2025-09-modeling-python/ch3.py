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
import matplotlib.pyplot as plt
import numpy as np

# %%
df = pl.DataFrame(dict(place=["olin", "wellesley"], num_of_bikes=[10, 2]))
df


# %%
def reassign_column(df, label, label_value, column, column_value):
    return df.with_columns(
        pl.when(df[label] == label_value)
        .then(column_value)
        .otherwise(df[column])
        .alias(column)
    )


reassign_column(df, "place", "olin", "num_of_bikes", 11)


# %%
def flip(p=0.5):
    return np.random.random() <= p


# %%
def bike_to_olin(df):
    if df.filter(df["place"] == "wellesley")["num_of_bikes"][0] == 0:
        return df

    df = reassign_column(
        df, "place", "wellesley", "num_of_bikes", df["num_of_bikes"] - 1
    )
    df = reassign_column(df, "place", "olin", "num_of_bikes", df["num_of_bikes"] + 1)
    return df


def bike_to_wellesley(df):
    if df.filter(df["place"] == "olin")["num_of_bikes"][0] == 0:
        return df

    df = reassign_column(df, "place", "olin", "num_of_bikes", df["num_of_bikes"] - 1)
    df = reassign_column(
        df, "place", "wellesley", "num_of_bikes", df["num_of_bikes"] + 1
    )
    return df


bike_to_olin(df)


# %%
def step(df, p1, p2):
    if flip(p1):
        df = bike_to_olin(df)

    if flip(p2):
        df = bike_to_wellesley(df)

    return df


# %%
def run(nsteps):
    timeseries = pl.DataFrame(
        schema=dict(time=pl.Int64, bikes_at_olin=pl.Int64, bikes_at_wellesley=pl.Int64)
    )
    df = pl.DataFrame(dict(place=["olin", "wellesley"], num_of_bikes=[10, 2]))
    for i in range(nsteps):
        df = step(df, 0.5, 0.33)
        bikes_at_olin = df.filter(pl.col("place") == "olin")["num_of_bikes"][0]
        bikes_at_wellesley = df.filter(pl.col("place") == "wellesley")["num_of_bikes"][
            0
        ]
        new_row = pl.DataFrame(
            dict(
                time=[i],
                bikes_at_olin=[bikes_at_olin],
                bikes_at_wellesley=[bikes_at_wellesley],
            )
        )
        timeseries = pl.concat([timeseries, new_row])

    plt.plot(
        timeseries["time"].to_numpy(),
        timeseries["bikes_at_olin"].to_numpy(),
        timeseries["bikes_at_wellesley"].to_numpy(),
    )
    plt.show()


run(10)
