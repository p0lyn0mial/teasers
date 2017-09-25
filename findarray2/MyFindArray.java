package findarray2;
// Teaser:
 // Implement a method that given two arrays as parameters will find the starting index where the second parameter occurs as a sub-array in the array given as the first parameter.
 // If given sub-array (second parameter) occurs more than once, then the method should return the starting index of the last occurrence
 // Your implementations should return -1 if the sub-array cannot be found.
 // Your implementation must implement the FindArray interface.
 // For extra points: implement it in an efficient way for large input arrays.
 //
 // Sample Input:
 // [4,9,3,7,8] and [3,7] should return 2.
 // [1,3,5] and [1] should return 0.
 // [7,8,9] and [8,9,10] should return -1.
 // [4,9,3,7,8,3,7,1] and [3,7] should return 5.
public class MyFindArray implements FindArray {
    public int findArray(int[] array, int[] subArray) {
        if (subArray.length > array.length) { return -1; }
        if (array.length == 0 || subArray.length == 0) { return -1; }
   
        int skip = subArray.length;

        // iterate the array backwards, each time skipping some elements
        for (int i = array.length - skip; i >= 0; i = i - skip) {
            // if there is a match on the first index,
            // check the remaining positions
            if (array[i] == subArray[0]) {    
                boolean foundSolution = true;

                for (int j = 1; j < subArray.length; j++) {
                    if (array[i+j] != subArray[j]) {  
                        foundSolution = false;
                        break;
                    }
               }

               if (foundSolution) { return i; }
               skip = subArray.length;
               continue;
          }

          // if there is match on the second value skip only be one position
          // the idea is to check whether numbers from the input array
          // overlap with the numbers from the sub-array on any positions
          // in that case we need to recalculate our skip.
          int newSkip = 0;
          boolean potentialSolution = false;
          int nextArrayPosition = 0;
          for (int j = 0; j < subArray.length; j++) {
              if (array[i+nextArrayPosition] != subArray[j]) {
                  if (potentialSolution) {          
                      potentialSolution = false;
                      break;
                  }
                  newSkip = newSkip + 1;
              } else {
                  potentialSolution = true;
                  nextArrayPosition = nextArrayPosition + 1;
              }
          }
 
          if (potentialSolution) { skip = newSkip;} 
	  else { skip = subArray.length; }
        }
	    
        return -1; 
    }

    public static void main(String[] args) {
	// scenario 1
	{
          int[] input = { 8,3,7 };
          int[] output = { 3,7 };
	  int expectedSolution = 1;

          MyFindArray mfa = new MyFindArray();
          int actualSolution = mfa.findArray(input, output);
	  if (actualSolution != expectedSolution) {
            System.out.format("scenario 1 failed, expected %d, got %d", expectedSolution, actualSolution);
            System.exit(-1);
	  }
	}
    }
}



