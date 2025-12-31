toDigits :: Integer -> [Integer]
toDigits n = go n []
  where
    go n xs
      | n <= 0 = xs
      | otherwise = go (n `div` 10) (n `mod` 10 : xs)

toDigitsRev :: Integer -> [Integer]
toDigitsRev n = reverse (toDigits n)

-- Method 1
-- doubleEveryOther :: [Integer] -> [Integer]

-- doubleEveryOther n = reverse (zipWith f [0 .. length n - 1] (reverse n))
--   where
--     f index value
--       | even index = value
--       | otherwise = 2 * value

-- Method 2
doubleEveryOther :: [Integer] -> [Integer]
doubleEveryOther a = zipWith f [0 .. length a - 1] a
  where
    lenEven = even (length a)
    f index val
      | (lenEven && even index) || (not lenEven && odd index) = 2 * val
      | otherwise = val

sumDigits :: [Integer] -> Integer
sumDigits a = sum (map (sum . toDigits) a)

validate :: Integer -> Bool
validate x = mod 10 (sumDigits $ doubleEveryOther $ toDigits x) == 0
