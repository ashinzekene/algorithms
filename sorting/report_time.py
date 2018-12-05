import time

def report_time(func, *args):
    st = time.time()
    for _ in range(9999):
        func(*args)
    total = round(time.time() - st, 5)
    print('-------------------------------------')
    print('{} took {} seconds'.format(func.__name__, total))
    return func(*args)
