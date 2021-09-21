import os

# Directory to fix
DIR = './logs/'

def main():
    filenames = os.listdir(DIR)
    
    if not len(filenames): return
    filenames.sort()
    current = filenames[0][:10]

    i = 0
    while(i < len(filenames)):
        # Get all files with the same date
        workingFiles = []
        while(i < len(filenames) and filenames[i][:10] == current):
            workingFiles.append((int(filenames[i].split('-')[3][:-4]), filenames[i]))
            i += 1

        # Sort by the pre-existing number
        workingFiles.sort()

        # Check fof new date
        if i < len(filenames):
            if filenames[i][:10] != current:
                current = filenames[i][:10]

        # Write file with new number and .temp
        tempFiles = []
        for j, f in enumerate(workingFiles):
            newName = DIR + f[1][:10] + "-" + str(len(workingFiles)-j) + '.log.temp'
            print("RENAMING", DIR+f[1], "->", newName)
            os.rename(DIR+f[1], newName)
            tempFiles.append(newName)

        # Remove .temp
        for f in tempFiles:
            os.rename(f, f[:len(f)-5])

if __name__ == '__main__':
    main()