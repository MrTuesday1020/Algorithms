import heapq
import math

graph={
	"A":{"B":5,"C":1},
	"B":{"A":5,"C":2,"D":1},
	"C":{"A":1,"B":2,"D":4,"E":8},
	"D":{"B":1,"C":4,"E":3,"F":6},
	"E":{"C":8,"D":3},
	"F":{"D":6}
}


def init_distance(graph, start):
	'''init all the distances as infinite'''
	distance = {start:0}
	for vertex in graph:
		if vertex != start:
			distance[vertex] = math.inf
	return distance
	

def dijkstra(graph, start):
	'''Do Breadth First Searching over a graph with a particular starting point'''
	# use a priority queue
	pqueue = []
	heapq.heappush(pqueue, (0, start))
	# record the nodes which already exist
	seen = set()
	
	parents = {start:None}
	distance = init_distance(graph, start)
	
	while(len(pqueue) > 0):
		pair = heapq.heappop(pqueue)
		dist = pair[0]
		vertex = pair[1]
		seen.add(vertex)
		
		nodes = graph[vertex].keys()
		for node in nodes:
			if node not in seen:
				if dist + graph[vertex][node] < distance[node]:
					# update priority queue
					heapq.heappush(pqueue, (dist + graph[vertex][node], node))
					# update the parent of the current node
					parents[node] = vertex
					# update the distance from the root to the current node
					distance[node] = dist + graph[vertex][node]
	
	return parents, distance

parents, distance = dijkstra(graph, "A")
print(parents)
print(distance)