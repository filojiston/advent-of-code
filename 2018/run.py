import os
import sys

if not len(sys.argv) == 2:
    print("Usage: python3 run.py [day-to-run]")
    exit(1)

day_to_run = sys.argv[1]
if sys.platform == "linux" or sys.platform == "linux2":
    os.system(f"python3 -B -m {day_to_run}.app")
elif sys.platform == "darwin":
    os.system(f"python3 -B -m {day_to_run}.app")
elif sys.platform == "win32":
    os.system(f"python3.12.exe -B -m {day_to_run}.app")
