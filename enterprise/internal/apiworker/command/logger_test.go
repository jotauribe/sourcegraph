package command

// func TestLogger(t *testing.T) {
// 	logger := NewLogger("baz2", "BAR2")

// 	for i := 0; i < 3; i++ {
// 		outReader, outWriter := io.Pipe()
// 		errReader, errWriter := io.Pipe()

// 		go func() {
// 			defer outWriter.Close()
// 			defer errWriter.Close()

// 			_, _ = io.Copy(outWriter, bytes.NewReader([]byte(fmt.Sprintf("foo%[1]d bar%[1]d\nbaz%[1]d", i+1))))
// 			_, _ = io.Copy(errWriter, bytes.NewReader([]byte(fmt.Sprintf("FOO%[1]d BAR%[1]d\nBAZ%[1]d", i+1))))
// 		}()

// 		logger.RecordCommand(
// 			[]string{"test", strconv.FormatInt(int64(i)+1, 10)},
// 			outReader,
// 			errReader,
// 		)
// 	}

// 	entries := logger.Entries()
// 	if len(entries) != 3 {
// 		t.Fatalf("unexpected number of entries. want=%d have=%d", 3, len(entries))
// 	}

// 	expected1 := `
// 		stderr: BAZ1
// 		stderr: FOO1 BAR1
// 		stdout: baz1
// 		stdout: foo1 bar1
// 		`
// 	if diff := cmp.Diff([]string{"test", "1"}, entries[0].Command); diff != "" {
// 		t.Errorf("unexpected command (-want +got):\n%s", diff)
// 	}
// 	if diff := cmp.Diff(normalizeLogs(expected1), normalizeLogs(entries[0].Out)); diff != "" {
// 		t.Errorf("unexpected log output (-want +got):\n%s", diff)
// 	}

// 	expected2 := `
// 		stderr: BAZ2
// 		stderr: FOO2 ******
// 		stdout: ******
// 		stdout: foo2 bar2
// 		`
// 	if diff := cmp.Diff([]string{"test", "2"}, entries[1].Command); diff != "" {
// 		t.Errorf("unexpected command (-want +got):\n%s", diff)
// 	}
// 	if diff := cmp.Diff(normalizeLogs(expected2), normalizeLogs(entries[1].Out)); diff != "" {
// 		t.Errorf("unexpected log output (-want +got):\n%s", diff)
// 	}

// 	expected3 := `
// 		stderr: BAZ3
// 		stderr: FOO3 BAR3
// 		stdout: baz3
// 		stdout: foo3 bar3
// 		`
// 	if diff := cmp.Diff([]string{"test", "3"}, entries[2].Command); diff != "" {
// 		t.Errorf("unexpected command (-want +got):\n%s", diff)
// 	}
// 	if diff := cmp.Diff(normalizeLogs(expected3), normalizeLogs(entries[2].Out)); diff != "" {
// 		t.Errorf("unexpected log output (-want +got):\n%s", diff)
// 	}
// }

// func normalizeLogs(text string) (filtered []string) {
// 	for _, line := range strings.Split(text, "\n") {
// 		if line := strings.TrimSpace(line); line != "" {
// 			filtered = append(filtered, line)
// 		}
// 	}
// 	sort.Strings(filtered)

// 	return filtered
// }
