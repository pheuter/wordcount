(ns wordcount)

(defn word-count
  [s]
  (sort-by last >
           (reduce
             #(assoc %1 %2 (inc (%1 %2 0)))
             {}
             (re-seq #"\w+" s))))
