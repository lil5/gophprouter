<?php

// Functions
// ################

function showErrors(): void {
	ini_set('display_errors', 1);
	ini_set('display_startup_errors', 1);
	error_reporting(E_ALL);
}

function validUrl($s): bool {
	return preg_match('/[^a-zA-Z0-9\-_\/%]/', $s) !== 1;
}

function validMethod(string $method): bool {
	switch ($method) {
		case 'DELETE':
		case 'GET':
		case 'HEAD':
		case 'OPTIONS':
		case 'PATCH':
		case 'POST':
		case 'PUT':
			return true;
			break;
	
		default:
			return false;
			break;
	}
}

function returnError(int $code): void {
	header('Content-Type: application/json');
	echo('{"ok": false, "status": '.$code.'}');
	http_response_code($code);
	exit;
}

function getRequestBody(): string {
	$rawdata = file_get_contents('php://input');

	return urlencode($rawdata);
}

// Variables
// ################

$url = $_SERVER['REQUEST_URI'];
$method = $_SERVER['REQUEST_METHOD'];
$body = getRequestBody();

// Main
// ################

showErrors();

if (!validMethod($method)) {
	returnError(400);
	exit;
}

if ($method === 'OPTIONS') {
	header('Access-Control-Allow-Origin: *');
	http_response_code(204);
	exit;
}

if (!validUrl($url)) {
	returnError(400);
	exit;
}

$command='./main -path="'.$url.'" -method="'.$method.'" -body="'.$body.'"';
$return_var=null;

passthru($command, $return_var);

if($return_var !== 0) {
	returnError(500);
	exit;
}

http_response_code(200);

