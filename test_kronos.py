# Let's come back to testing later

import os
import subprocess
import time
import pytest

# Set the path to the executable binary
KRONOS_PATH = "./main"

# Define the paths for the test CSV files
CSV_PATH_VALID = "test_valid.csv"
CSV_PATH_MISSING_FIELD = "test_missing_field.csv"
CSV_PATH_EMPTY = "test_empty.csv"
CSV_PATH_BAD_REQUEST = "test_bad_request.csv"
CSV_PATH_BAD_HEADERS = "test_bad_headers.csv"
CSV_PATH_BAD_METHOD = "test_bad_method.csv"
CSV_PATH_BAD_REQUEST_BODY = "test_bad_request_body.csv"
CSV_PATH_MALFORMED_HEADERS = "test_malformed_headers.csv"

def test_kronos_csv_valid():
    # Test that running Kronos with a valid CSV file works correctly.
    threshold = 5
    output = subprocess.check_output([KRONOS_PATH, "-csv", CSV_PATH_VALID, "-threshold", str(threshold)]).decode("utf-8")
    assert "Test passed" in output

def test_kronos_csv_invalid_threshold():
    # Test that running Kronos with an invalid threshold value falls back to the default.
    threshold = "invalid_threshold"
    output = subprocess.check_output([KRONOS_PATH, "-csv", CSV_PATH_VALID, "-threshold", str(threshold)]).decode("utf-8")
    assert "Invalid threshold value, using default value of 1 seconds" in output

def test_kronos_csv_response_time_ok():
    # Test that running Kronos with a CSV file where the response time is within the threshold passes.
    threshold = 5
    output = subprocess.check_output([KRONOS_PATH, "-csv", CSV_PATH_VALID, "-threshold", str(threshold)]).decode("utf-8")
    assert "Test passed" in output

def test_kronos_csv_invalid_path():
    # Test that running Kronos with an invalid CSV file path fails.
    threshold = 5
    with pytest.raises(subprocess.CalledProcessError):
        subprocess.check_output([KRONOS_PATH, "-csv", "invalid_path.csv", "-threshold", str(threshold)])

def test_kronos_csv_missing_field():
    # Test that running Kronos with a CSV file that is missing a field fails.
    threshold = 5
    with pytest.raises(subprocess.CalledProcessError):
        subprocess.check_output([KRONOS_PATH, "-csv", CSV_PATH_MISSING_FIELD, "-threshold", str(threshold)])

def test_kronos_csv_empty():
    # Test that running Kronos with an empty CSV file fails.
    threshold = 5
    with pytest.raises(subprocess.CalledProcessError):
        subprocess.check_output([KRONOS_PATH, "-csv", CSV_PATH_EMPTY, "-threshold", str(threshold)])

def test_kronos_csv_bad_request():
    # Test that running Kronos with a CSV file containing an invalid URL returns an error.
    threshold = 5
    with pytest.raises(subprocess.CalledProcessError):
        subprocess.check_output([KRONOS_PATH, "-csv", CSV_PATH_BAD_REQUEST, "-threshold", str(threshold)])

def test_kronos_csv_bad_headers():
    # Test that running Kronos with a CSV file containing bad headers returns an error.
    threshold = 5
    with pytest.raises(subprocess.CalledProcessError):
        subprocess.check_output([KRONOS_PATH, "-csv", CSV_PATH_BAD_HEADERS, "-threshold", str(threshold)])

def test_kronos_csv_bad_method():
    # Test that running Kronos with a CSV file containing an invalid HTTP method returns an error.
    csv_path = "api_info_bad_method.csv"
    threshold = 5
    with pytest.raises(subprocess.CalledProcessError) as excinfo:
        subprocess.check_output([KRONOS_PATH, "-csv", csv_path, "-threshold", str(threshold)])
    assert "Error parsing CSV data" in str(excinfo.value)
