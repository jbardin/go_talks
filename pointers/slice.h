// runtime.h
struct  Slice
{
    byte*   array;      // actual data
    uintgo  len;        // number of elements
    uintgo  cap;        // allocated number of elements
};

