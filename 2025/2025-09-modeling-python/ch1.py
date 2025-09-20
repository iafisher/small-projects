# ---
# jupyter:
#   jupytext:
#     formats: ipynb,py:percent
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
a = 9.8  # earth's gravity

# %%
t = 3.4  # arbitrary time to measure

# %%
v = a * t

# %%
v

# %%
x = a * t**2 / 2

# %%
x  # how far has the penny fallen?

# %%
h = 381  # height of the Empire State building, in meters

# %%
from numpy import sqrt

t = sqrt(2 * h / a)
t  # seconds until it hits the ground

# %%
v = a * t
v  # velocity (m/s) when it hits the ground

# %%
from pint import UnitRegistry

units = UnitRegistry()

# %%
a = 9.8 * units.meter / units.second**2
a

# %%
t = 3.4 * units.seconds

# %%
a * t**2 / 2

# %%
v = a * t
v.to(units.mile / units.hour)

# %%
h = 381 * units.meter

# %%
pole_height = 10 * units.foot
print(pole_height + h, h + pole_height)

# %%
a + t

# %%
(381 * unit.meter) / (29 * unit.meter / unit.second)

# %%
# exercise 1.6
m = unit.meter
s = unit.second

a = 9.8 * m / s**2
h = 381 * m
terminal_velocity = 29 * m / s

# How long to reach 29 m/s?
# v = a * t
# 29 m/s = a * t
# t = (29 m/s) / a

t_until_terminal = terminal_velocity / a

# How far did it fall in that time?
# y(t) = h - (a * t**2 / 2)

height_at_terminal = h - (a * t_until_terminal**2 / 2)
height_at_terminal

# %%
# How long to reach ground once at terminal velocity?
# y(t) = height_at_terminal - 29 * t
# 0 = height_at_terminal - 29 * t
# t = height_at_terminal / 29
t_from_terminal_until_ground = height_at_terminal / (29 * m / s)
t_from_terminal_until_ground

# %%
t_until_terminal + t_from_terminal_until_ground
