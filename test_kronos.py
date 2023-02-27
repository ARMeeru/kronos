import csv
import os
import subprocess
import time
import pytest

# Set the path to the Kronos executable
KRONOS_PATH = "./kronos"

def test_kronos_csv():
    # Test that running Kronos with a valid CSV file works correctly.
    csv_path = "api_info.csv"
    threshold = 5
    output = subprocess.check_output([KRONOS_PATH, "-csv", csv_path, "-threshold", str(threshold)]).decode("utf-8")
    assert "Test passed" in output

def test_kronos_csv_invalid_threshold():
    # Test that running Kronos with an invalid threshold value falls back to the default.
    csv_path = "api_info.csv"
    threshold = "invalid_threshold"
    output = subprocess.check_output([KRONOS_PATH, "-csv", csv_path, "-threshold", str(threshold)]).decode("utf-8")
    assert "Invalid threshold value, using default value of 1 seconds" in output

def test_kronos_csv_response_time_ok():
    # Test that running Kronos with a CSV file where the response time is within the threshold passes.
    csv_path = "api_info.csv"
    threshold = 5
    output = subprocess.check_output([KRONOS_PATH, "-csv", csv_path, "-threshold", str(threshold)]).decode("utf-8")
    assert "Test passed" in output

# def test_kronos_csv_invalid_path():
#     # Test that running Kronos with an invalid CSV file path fails.
#     csv_path = "invalid_path.csv"
#     threshold = 5
#     with pytest.raises(subprocess.CalledProcessError):
#         subprocess.check_output([KRONOS_PATH, "-csv", csv_path, "-threshold", str(threshold)])

# def test_kronos_csv_missing_field():
#     # Test that running Kronos with a CSV file that is missing a field fails.
#     csv_path = "api_info_missing_field.csv"
#     threshold = 5
#     with pytest.raises(subprocess.CalledProcessError):
#         subprocess.check_output([KRONOS_PATH, "-csv", csv_path, "-threshold", str(threshold)])

# def test_kronos_csv_empty():
#     # Test that running Kronos with an empty CSV file fails.
#     csv_path = "api_info_empty.csv"
#     threshold = 5
#     with pytest.raises(subprocess.CalledProcessError):
#         subprocess.check_output([KRONOS_PATH, "-csv", csv_path, "-threshold", str(threshold)])

# def test_kronos_csv_bad_request():
#     # Test that running Kronos with a CSV file containing an invalid URL returns an error.
#     csv_path = "api_info_bad_request.csv"
#     threshold = 5
#     with pytest.raises(subprocess.CalledProcessError):
#         subprocess.check_output([KRONOS_PATH, "-csv", csv_path, "-threshold", str(threshold)])

# def test_kronos_csv_bad_headers():
#     # Test that running Kronos with a CSV file containing bad headers returns an error.
#     csv_path = "api_info_bad_headers.csv"
#     threshold = 5
#     with pytest.raises(subprocess.CalledProcessError):
#         subprocess.check_output([KRONOS_PATH, "-csv", csv_path, "-threshold", str(threshold)])

# def test_kronos_csv_bad_method():
#     # Test that running Kronos with a CSV file containing an invalid HTTP method returns an error.
#     csv_path = "api_info_bad_method.csv"
#     threshold = 5
#     with pytest.raises(subprocess.CalledProcessError):
#         subprocess.check_output([KRONOS_PATH, "-csv", csv_path, "-threshold", str(threshold)])

# def test_kronos_csv_bad_request_body():
#     # Test that running Kronos with a CSV file containing an invalid request body returns an error.
#     csv_path = "api_info_bad_request_body.csv"
#     threshold = 5
#     with pytest.raises(subprocess.CalledProcessError):
#         subprocess.check_output([KRONOS_PATH, "-csv", csv_path, "-threshold", str(threshold)])

# def test_kronos_csv_malformed_headers():
#     # Test that running Kronos with a CSV file containing malformed headers returns an error.
#     csv_path = "api_info_malformed_headers.csv"
#     threshold = 5
#     with pytest.raises(subprocess.CalledProcessError):
#         subprocess.check_output([KRONOS_PATH, "-csv", csv_path, "-threshold", str(threshold)])
