graph={
	"A":["B","C"],
	"B":["A","C","D"],
	"C":["A","B","D","E"],
	"D":["B","C","E","F"],
	"E":["C","D"],
	"F":["D"]
}

def BFS(graph, start):
	'''Do Breadth First Searching over a graph with a particular starting point'''
	# use a queue: first in first out
	queue = []
	queue.append(start)
	seen = set()
	seen.add(start)
	parents = {start:None}
	
	result = []
	while(len(queue) > 0):
		vertex = queue.pop(0)
		nodes = graph[vertex]
		for node in nodes:
			if node not in seen:
				queue.append(node)
				seen.add(node)
				parents[node] = vertex
		result.append(vertex)
	
	return result, parents

def find_path(parents, node):
	'''find a path from a given node to the root'''
	path = []
	while node:
		path.append(node)
		node = parents[node]
	return path[::-1]

	
result, parents = BFS(graph, "E")
print(result)

path = find_path(parents, "B")
print(path)