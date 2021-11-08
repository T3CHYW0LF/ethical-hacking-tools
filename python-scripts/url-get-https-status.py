import requests
fname=sys.argv[1]
print(fname)
file= open(fname, 'r')
Lines = file.readlines()
for line in Lines:
    line = line.replace('\n','')
    res = requests.get(line)
    status_code=res.status_code
    if status_code == 200:
        print(f"ACTIVE-{status_code}:{line}")
    else:
        print(f"CHANGED-{status_code}:{line}")
