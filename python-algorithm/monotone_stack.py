def left_greater(nums: list[int]) -> list[int]:
    n = len(nums)
    left = [-1] * n
    st = []
    for i, x in enumerate(nums):
        while st and nums[st[-1]] <= x:
            st.pop()
        if st:
            left[i] = st[-1]
        st.append(i)
    return left


def right_greater(nums: list[int]) -> list[int]:
    n = len(nums)
    right = [-1] * n
    st = []
    for i, x in enumerate(nums):
        while st and nums[st[-1]] >= x:
            st.pop()
        if st:
            right[i] = st[-1]
        st.append(i)
    return right
