toEval=$(cut -f 1 <<< "$tasks")
tasks=$(cut -f 2- <<< "$tasks")
source /ipfs/$toEval