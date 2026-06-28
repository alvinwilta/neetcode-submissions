type LRUCache struct {
	cap int
    storage map[int]*Cache
	first *Cache
	last *Cache
}

type Cache struct {
	key int
	val int
	prev *Cache
	next *Cache
}

func Constructor(capacity int) LRUCache {
	first := &Cache{}
	last := &Cache{}
	first.next = last
	last.prev = first
    return LRUCache{
		cap: capacity,
		storage: make(map[int]*Cache),
		first: first,
		last: last,
	}
}

func (this *LRUCache) Get(key int) int {
	cache, exist := this.storage[key]
	if !exist {
		return -1
	}
	this.access(cache)
	return cache.val
}

func (this *LRUCache) Put(key int, value int) {
	cache, exist := this.storage[key]
	if !exist {
		// new element
		newCache := &Cache{
			key: key,
			val: value,
			prev: this.last,
		}
		if len(this.storage) >= this.cap {
			// over cap
			evictedCache := this.first.next
			this.first.next = this.first.next.next
			this.first.next.prev = this.first
			delete(this.storage,evictedCache.key)
		} 
		newCache.prev = this.last.prev
		this.last.prev.next = newCache
		this.last.prev = newCache
		newCache.next = this.last
		this.storage[key] = newCache
	} else {
		// update current
		cache.val = value
		this.access(cache)
	}
}

func (this *LRUCache) access(cache *Cache) {
	// cut off link
	cache.prev.next = cache.next
	cache.next.prev = cache.prev

	// move cache to end
	cache.next = this.last
	cache.prev = this.last.prev
	this.last.prev.next = cache
	this.last.prev = cache
}

