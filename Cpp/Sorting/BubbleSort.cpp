//******************************************************************************************************************************************
// Given an array of integers called nums,sort the array in non-decreasing order
// using the bubble sort algorithm and return the sorted array. A sorted array
// in non-decreasing order is an array where each element is greater than or
// equal to all preceding elements in the array. Examples: Input: nums = [7, 4,
// 1, 5, 3] Output: [1, 3, 4, 5, 7] Explanation: 1 <= 3 <= 4 <= 5 <= 7. Thus the
// array is sorted in non-decreasing order. Input: nums = [5, 4, 4, 1, 1]
// Output: [1, 1, 4, 4, 5]
// Explanation: 1 <= 1 <= 4 <= 4 <= 5.
// Thus the array is sorted in non-decreasing order.
//******************************************************************************************************************************************

#include <vector>
using namespace std;
class Solution {
public:
  vector<int> bubbleSort(vector<int> &nums) {
    for (size_t i = nums.size() - 1; i > 0; i--) {
      bool swapped = false;
      for (size_t j = 0; j < i; j++) {
        if (nums[j] > nums[j + 1]) {
          int temp = nums[j];
          nums[j] = nums[j + 1];
          nums[j + 1] = temp;
          swapped = true;
        }
      }
      if (!swapped) {
        break;
      }
    }
    return nums;
  }
};
