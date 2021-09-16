# qsub

extremely WIP - will solifidy quickly - goal is to mock qsub for testing grid execution without the grid

## Installation

Nothing yet

## Usage

one issue is that it does not 100% represent the exact statement as passed:

```
./qsub -cwd -l pmem="16gb" run.sh  
```

```
{
        "Script": "run.sh",
        "Flags": {
                "cwd": "",
                "l": "pmem=16gb"
        }
}
```

for example, the specific quotes around pmem="16gb" is not captured


Why parse the flags? well this way we could more specifically target unit tests if we so choose to 
both check specific elements, and/or do input validation

```
./qsub -cwd -pe orte 8 run.sh  
```

Can directly do:

```
 flagSet.ParallelEnvironment.Value.Slots = 8
 ```

 vs the originating parsed flags

 ```
 "pe": "orte 8"
 ```

In addition, will do input validation:

```
./qsub -cwd -pe asdf 8 run.sh
```

```
could not set parallel environment `asdf 8`
2021/09/16 17:36:06 unknown parallel environment
```