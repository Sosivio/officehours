
import argparse
import time

GB = 1024 * 1024 * 1024
GB = 1024 * 1024 

def parse_cmd_args():
    """
    Parses command line args
    """
    parser = argparse.ArgumentParser(description='Memory Eating utility for python')
    parser.add_argument('-m', '--memory', type=int, default=-1, help='The amount of memory in gigs to eat', required = True)
    args = parser.parse_args()
    return args


def eat_memory(mem_to_eat):
    """
    Eats memeory
    :param mem_to_eat: The amount of memory to eat in gigs
    """
    global GB
    eat=""
    #eat = "a" * GB * mem_to_eat
    for x in range(mem_to_eat):
        eat = eat+("a" * GB)
        time.sleep(0.2)
    while True:
        time.sleep(1)
        


def main():
    """
    Main sentinel
    """
    mem_to_eat = parse_cmd_args().memory
    eat_memory(mem_to_eat)

if __name__ == "__main__":
    main()
