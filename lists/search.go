package lists

func (ll *LinkedList) Search(value interface{}) *Iterator {
	cleanedVal := ll.valueType.Validate(value)
	for it := ll.Begin(); it != ll.End(); it = it.Next() {
		if cleanedVal.Key() == it.currentNode.Value.Key() {
			return it
		}
	}
	return nil
}

