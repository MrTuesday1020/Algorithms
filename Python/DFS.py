graph={
	"A":["B","C"],
	"B":["A","C","D"],
	"C":["A","B","D","E"],
	"D":["B","C","E","F"],
	"E":["C","D"],
	"F":["D"]
}

def DFS(graph, start):
	'''Do Depth First Searching over a graph with a particular starting point'''
	# use a stack: first in last out
	stack = []
	stack.append(start)
	seen = set()
	seen.add(start)
	
	result = []
	while(len(stack) > 0):
		vertex = stack.pop()
		nodes = graph[vertex]
		for node in nodes:
			if node not in seen:
				stack.append(node)
				seen.add(node)
		result.append(vertex)
	
	return result
	
print(DFS(graph, "A"))