# Rotate a Linked List

> Striver SDE Sheet 037 · Day 7 LL & Arrays — join the list into a ring, then cut it n−k nodes in; k mod n first, since rotating by the length is a no-op. Go, stdlib only, tested.

## Problem

Rotate a singly linked list right by k places — the last k nodes move, in order, to the front. k can be far larger than the list.

## Approach

Rotation never changes the cyclic order of nodes, only where the list "starts" — which suggests literally making it a cycle. One pass finds the tail and length n; connect tail to head to close the ring, reduce k modulo n (huge k collapses immediately), then walk n−k−1 steps from the old head to the new tail and cut the ring there. Two short walks, no node allocation, and the k ≥ n case costs nothing extra.

**Complexity:** O(n) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
