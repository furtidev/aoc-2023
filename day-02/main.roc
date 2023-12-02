app "aoc-day2"
    packages { pf: "https://github.com/roc-lang/basic-cli/releases/download/0.7.0/bkGby8jb0tmZYsy2hg1E_B2QrCgcSTxdUlHtETwm5m4.tar.br" }
    imports [
        pf.Stdout,
        "input.txt" as input : Str,
    ]
    provides [main] to pf

# We have 12 red, 13 green and 14 blue cubes.

main =
    Stdout.line "\(input |> part1)"

parseLine = \line ->
    # the plan is to map all items like this: [[ID, SUBSETS],...]. 
    identifier = List.get (Str.split ((List.get (Str.split line ": ") 0) |> Result.withDefault "") " ") 1 |> Result.withDefault ""
    games = (List.get (Str.split line ": ") 1) |> Result.withDefault ""
    [identifier, games]

subsetParser = \subset ->
    # the job of this function is to parse the subset string into a List <[13 red cubes, 5 blue cubes...]>, for the sake of simplicity for later when we eventually calculate the solution.
    splitSubset = Str.split subset "; "
    List.walk splitSubset [] \state, elem ->
        splitElem = Str.split elem ", "
        List.concat state splitElem

# 'game' represents individual lines in <input.txt>
evaluateGame = \state, game ->
    id = Str.toU32 (List.get game 0 |> Result.withDefault "0") |> Result.withDefault 0
    subset = subsetParser (List.get game 1 |> Result.withDefault "")

    invalidEntries = 
        List.walk subset 0 \invalidCount, item -> 
            splitItem = Str.split item " "

            countStr = List.get splitItem 0 |> Result.withDefault "0"
            count = Str.toU32 countStr |> Result.withDefault 0
            color = List.get splitItem 1 |> Result.withDefault ""

            when color is
                "red" ->
                    if count > 12 then
                        invalidCount + 1
                    else
                        invalidCount
                "green" ->
                    if count > 13 then
                        invalidCount + 1
                    else
                        invalidCount
                "blue" ->
                    if count > 14 then
                        invalidCount + 1
                    else
                        invalidCount
                _ -> invalidCount + 1

    if invalidEntries == 0 then
        state + id
    else
        state

part1 = \data ->
    lines = Str.split data "\n"
    parsedLines = List.map lines parseLine
    sum = List.walk parsedLines 0 evaluateGame
    dbg sum
    ""