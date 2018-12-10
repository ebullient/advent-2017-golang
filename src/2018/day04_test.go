package days

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"testing"
)

var wallFormat = regexp.MustCompile(`\[([0-9-]*) (\d+):(\d+)\] (.*)`)
var guardShift = regexp.MustCompile(`Guard #(\d+) begins shift`)

type WallRecord struct {
	sortString string
	minute     int
	text       string
}

func (w WallRecord) String() string {
	return fmt.Sprintf("%s: %d %s", w.sortString, w.minute, w.text)
}

type SleepLog struct {
	asleep  int
	minutes [60]int
	//	entries []WallRecord
}

var test4SampleInput = []string{
	"[1518-11-01 00:00] Guard #10 begins shift",
	"[1518-11-01 00:05] falls asleep",
	"[1518-11-01 00:25] wakes up",
	"[1518-11-01 00:30] falls asleep",
	"[1518-11-01 00:55] wakes up",
	"[1518-11-01 23:58] Guard #99 begins shift",
	"[1518-11-02 00:40] falls asleep",
	"[1518-11-02 00:50] wakes up",
	"[1518-11-03 00:05] Guard #10 begins shift",
	"[1518-11-03 00:24] falls asleep",
	"[1518-11-03 00:29] wakes up",
	"[1518-11-04 00:02] Guard #99 begins shift",
	"[1518-11-04 00:36] falls asleep",
	"[1518-11-04 00:46] wakes up",
	"[1518-11-05 00:03] Guard #99 begins shift",
	"[1518-11-05 00:45] falls asleep",
	"[1518-11-05 00:55] wakes up"}

var test4ScrambledInput = []string{
	"[1518-11-01 00:00] Guard #10 begins shift",
	"[1518-11-01 00:05] falls asleep",
	"[1518-11-02 00:50] wakes up",
	"[1518-11-04 00:46] wakes up",
	"[1518-11-05 00:03] Guard #99 begins shift",
	"[1518-11-05 00:45] falls asleep",
	"[1518-11-03 00:24] falls asleep",
	"[1518-11-03 00:29] wakes up",
	"[1518-11-04 00:02] Guard #99 begins shift",
	"[1518-11-04 00:36] falls asleep",
	"[1518-11-03 00:05] Guard #10 begins shift",
	"[1518-11-01 23:58] Guard #99 begins shift",
	"[1518-11-02 00:40] falls asleep",
	"[1518-11-01 00:25] wakes up",
	"[1518-11-01 00:30] falls asleep",
	"[1518-11-01 00:55] wakes up",
	"[1518-11-05 00:55] wakes up"}

func SortLog(logEntries []string) []WallRecord {
	records := []WallRecord{}

	for _, entry := range logEntries {
		match := wallFormat.FindStringSubmatch(entry)
		if match == nil {
			continue
		}
		min, e := strconv.Atoi(match[3])
		check(e)
		b := match[4][0]
		r := WallRecord{
			match[1] + match[2] + match[3] + string(b),
			min,
			match[4],
		}
		records = append(records, r)
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].sortString < records[j].sortString
	})

	return records
}

func GetID(text string) int {
	match := guardShift.FindStringSubmatch(text)
	i, e := strconv.Atoi(match[1])
	check(e)
	return i
}

func ParseRecords(records []WallRecord) (map[int]SleepLog, int, int, int) {
	logs := map[int]SleepLog{}

	var (
		id        int
		guard     SleepLog
		last      int
		mostSleep int
		sameTime  int
	)
	sleepiest := -1
	consistentTime := -1
	consistentGuard := -1

	for _, record := range records {
		if record.text[0] == 'G' {
			id = GetID(record.text)
			last = -1
		}
		guard = logs[id]
		//	guard.entries = append(guard.entries, record)

		switch record.text[0] {
		case 'f':
			last = record.minute
		case 'w':
			if last >= 0 {
				guard.asleep += record.minute - last
				for i := last; i < record.minute; i++ {
					guard.minutes[i]++
					if guard.minutes[i] > sameTime {
						sameTime = guard.minutes[i]
						consistentTime = i
						consistentGuard = id
					}
				}
				if guard.asleep > mostSleep {
					mostSleep = guard.asleep
					sleepiest = id
				}
				last = -1
			}
		}
		logs[id] = guard
	}

	return logs, sleepiest, consistentTime, consistentGuard
}

func FindBestMinute(log SleepLog) int {
	max := -1
	best := -1

	for i, v := range log.minutes {
		if v > max {
			max = v
			best = i
		}
	}

	return best
}

func TestSampleData_4part1(t *testing.T) {
	sorted := SortLog(test4SampleInput)
	scrambled := SortLog(test4ScrambledInput)
	if !reflect.DeepEqual(sorted, scrambled) {
		t.Error("Bad sort! Bad!")
	}

	parsed, sleepiest, consistentTime, consistentGuard := ParseRecords(sorted)
	if sleepiest != 10 {
		t.Error("Expected Guard #10 to be sleepiest, got", sleepiest)
	}
	if consistentGuard != 99 {
		t.Error("Expected Guard #99 to be most consistently asleep, got", consistentGuard)
	}
	if consistentTime != 45 {
		t.Error("Expected minute 45 as most consistently asleep, got", consistentTime)
	}

	guard10 := parsed[10]
	best := FindBestMinute(guard10)
	if guard10.asleep != 50 {
		t.Error("Expected Guard #10 to be asleep for 50 minutes, got", guard10.asleep)
	}
	if best != 24 {
		t.Error("Expected minute 24, got", best)
	}

	guard99 := parsed[99]
	if guard99.asleep != 30 {
		t.Error("Expected Guard #99 to be asleep for 30 minutes, got", guard99.asleep)
	}
}

func TestInput_4(t *testing.T) {
	content, err := ioutil.ReadFile("day04_input.txt")
	check(err)
	wallWriting := strings.Split(string(content), "\n")

	defer elapsed("TestInput_4")() // time execution of the rest

	records := SortLog(wallWriting)
	parsed, sleepiest, consistentGuard, consistentTime := ParseRecords(records)
	best := FindBestMinute(parsed[sleepiest])
	//fmt.Println(sleepiest, parsed[sleepiest], best)

	fmt.Println("Day 4 / Part 1 Result", sleepiest*best)
	fmt.Println("Day 4 / Part 2 Result", consistentGuard*consistentTime)
}
