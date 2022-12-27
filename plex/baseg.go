package plex

const plex_base_cc = `
template<typename T> struct dyn_array{
	T* val;
	inline dyn_array(size_t size): val(alloca(size * sizeof(T))){}
	dyn_array(dyn_array&) = delete;
}
`
