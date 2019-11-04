import math, random
from operator import itemgetter
from prettytable import PrettyTable

ROWCOUNT = 50
pCrossover = 0.8
pMutation = 0.1
iterations = 70000

def crossover(mat):
    # sortemat = sorted(mat, key=itemgetter(10), reverse=True)
    cross_pos = int(round(random.random()*6 + 1))
    cross_count = int(ROWCOUNT * pCrossover)
    pts = randperm(list(range(cross_count)))
        
    for i in range(int(cross_count/2)):
        parent1 = mat[pts[2*i - 1]][0:8]
        parent2 = mat[pts[2*i]][0:8]
        child1 = parent1[:cross_pos] + parent2[cross_pos:]
        child2 = parent2[:cross_pos] + parent1[cross_pos:]
        eval_expand(child1)
        mat.append(child1)
        eval_expand(child2)
        mat.append(child2)
    
def randperm(a):
    if(not a):
        return a
    b = []
    while(a.__len__()):
        r = random.choice(a)
        b.append(r)
        a.remove(r)
    return b
        
def printM(m):
    p = PrettyTable()
    for row in m:
        p.add_row(row)
    print(p.get_string(header=False, border=True))
    
def populate(matrix):
    for i in range(ROWCOUNT):  #rows
        row = []
        for j in range(8):  #columns
            if random.random() > 0.5:
                num=1
            else:
                num=0
            row.append(num)
        eval_expand(row)
        matrix.append(row)
    

#expect binary values only
def eval_expand(row):
    row.append(decimal(row))
    row.append(row[-1]/255.0*2*math.pi - math.pi)
    row.append(objectiveFunc(row[-1]))
    return row
    
def decimal(ro):
    sum=0
    pow=len(ro)-1
    for i in ro:
        if i==1:
            sum=sum+2**pow
        pow=pow-1    
    return sum
    
def objectiveFunc(val):
    return math.sin(val)
    
def mutate(mat):
    for i in range(int(ROWCOUNT * (1+pCrossover))):
        for j in range(8):  #unchanged, 'coz 8 bits to represent numbers
            if random.random() > (1.0-pMutation):    
                if(mat[i][j]):
                    mat[i][j]=0
                else:
                    mat[i][j]=1
                #print "Supposed to mutate matrix at ",i,j,"here, whose binary is ",mat[i][:8]
                #mutation done, now re-eval the thing
                newRow = eval_expand(mat[i][:8])
                mat[i] = newRow


if __name__=='__main__':
    print(f"Parameters: ROWCOUNT = {ROWCOUNT}, pCrossover = {pCrossover}, pMutation = {pMutation}, iterations = {iterations}")

    matrix=[]
    populate(matrix)
    print("\n***INITIAL STATE***\n")
    printM(matrix)
    
    for k in range(iterations):
        crossover(matrix)
        mutate(matrix)
        #kill
        matrix = sorted(matrix, key=itemgetter(10), reverse=True) #10 refers to column number
        matrix = matrix[:ROWCOUNT]
        
    print("***RESULT***\n")     
    printM(matrix)
