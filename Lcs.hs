{-# LANGUAGE BangPatterns #-}

import System.IO (getLine, print)
import Data.List (map, foldl')

lcs :: String -> String -> Int
lcs !a !b = last $ foldl' (calcLcsRow b) (map (const 0) b) a
  where
    calcLcsRow :: String -> [Int] -> Char -> [Int]
    calcLcsRow str !upRow !char = doTheStuff str 0 upRow 0
      where
        doTheStuff :: String -> Int -> [Int] -> Int -> [Int]
        doTheStuff _ _ [] _ = []
        doTheStuff (otherChar:string) !upLeft (up:prevRowRest) !left =
          let value = if char == otherChar then upLeft + 1 else max up left
          in value:doTheStuff string up prevRowRest value


main = lcs <$> getLine <*> getLine >>= print
