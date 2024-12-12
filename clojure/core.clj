(ns cloj.core
  (:require [clojure.set :as set]))

(defrecord Node [id data])
(defrecord Edge [from to])

(defn create-graph []
  "Create an empty directed graph."
  {:nodes {} :edges #{}})

(defn add-node [graph data]
  "Add a node with the given data to the graph. Returns the updated graph and the created node."
  (let [new-id (java.util.UUID/randomUUID)
        new-node (->Node new-id data)]
    [(assoc-in graph [:nodes new-id] new-node) new-node]))

(defn add-edge [graph from to]
  "Add a directed edge from one node to another. Returns the updated graph."
  (if (and (contains? (:nodes graph) (:id from))
           (contains? (:nodes graph) (:id to)))
    (update graph :edges conj (->Edge (:id from) (:id to)))
    (throw (ex-info "One or both nodes do not exist in the graph" {:from from :to to}))))

(defn get-neighbors [graph node]
  "Get a list of neighbors (nodes) for a given node."
  (let [node-ids (->> (:edges graph)
                      (filter #(= (:from %) (:id node)))
                      (map :to))]
    (map (fn [id] (get-in graph [:nodes id])) node-ids)))

(defn print-graph [graph]
  "Print the graph's nodes and edges in a readable format."
  (println "Nodes:")
  (doseq [[_ node] (:nodes graph)]
    (println (:id node) "->" (:data node)))
  (println "\nEdges:")
  (doseq [edge (:edges graph)]
    (println (:id (get-in graph [:nodes (:from edge)])) "->" (:id (get-in graph [:nodes (:to edge)])))))

;; Example usage
(let [graph (create-graph)
      [graph alice] (add-node graph "Alice")
      [graph bob] (add-node graph "Bob")
      [graph charlie] (add-node graph "Charlie")
      graph (add-edge graph alice bob)
      graph (add-edge graph bob charlie)
      graph (add-edge graph alice charlie)]
  (println "Neighbors of Alice:")
  (doseq [neighbor (get-neighbors graph alice)]
    (println (:data neighbor)))
  (println "\nGraph structure:")
  (print-graph graph))

