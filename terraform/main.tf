provider "aws" {
  region = "eu-central-1"
}

# S3 bucket for website hosting (kept private)
resource "aws_s3_bucket" "skatmate_web_bucket" {
  bucket = "skatmate-web-bucket"
}

# S3 bucket website configuration
resource "aws_s3_bucket_website_configuration" "website_config" {
  bucket = aws_s3_bucket.skatmate_web_bucket.id

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "error.html"
  }
}

# Lambda function to process requests
resource "aws_lambda_function" "data_handler" {
  function_name = "data_handler"
  runtime       = "java21"
  handler       = "de.polkow.skatmate.SpielarchivHandler::handleRequest"

  # Lambda code packaged as a .jar file
  filename         = "skatmate.jar"
  source_code_hash = filebase64sha256("skatmate.jar")

  role = aws_iam_role.lambda_exec.arn
}

# IAM role and policy for Lambda execution
resource "aws_iam_role" "lambda_exec" {
  name = "lambda_exec_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Action    = "sts:AssumeRole",
        Effect    = "Allow",
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_policy_attach" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

# API Gateway to trigger the Lambda function
resource "aws_apigatewayv2_api" "api_gateway" {
  name          = "skatmate_api"
  protocol_type = "HTTP"
}

resource "aws_lambda_permission" "api_gateway_permission" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.data_handler.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.api_gateway.execution_arn}/*"
}

resource "aws_apigatewayv2_integration" "api_integration" {
  api_id                  = aws_apigatewayv2_api.api_gateway.id
  integration_type        = "AWS_PROXY"
  integration_uri         = aws_lambda_function.data_handler.invoke_arn
  payload_format_version  = "2.0"
}

resource "aws_apigatewayv2_route" "api_route" {
  api_id    = aws_apigatewayv2_api.api_gateway.id
  route_key = "GET /data"  # Define your route path
  target    = "integrations/${aws_apigatewayv2_integration.api_integration.id}"
}

resource "aws_apigatewayv2_stage" "api_stage" {
  api_id      = aws_apigatewayv2_api.api_gateway.id
  name        = "default"
  auto_deploy = true
}

# CloudFront distribution to serve the S3 content
resource "aws_cloudfront_distribution" "cdn" {
  origin {
    domain_name = aws_s3_bucket.skatmate_web_bucket.bucket_regional_domain_name
    origin_id   = "s3-skatmate-web-bucket"

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.origin_access_identity.cloudfront_access_identity_path
    }
  }

  enabled             = true
  is_ipv6_enabled     = true
  default_root_object = "index.html"

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "s3-skatmate-web-bucket"
    viewer_protocol_policy = "redirect-to-https"

    forwarded_values {
      query_string = false
      cookies {
        forward = "none"
      }
    }
  }

  price_class = "PriceClass_100"

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }
}

# CloudFront Origin Access Identity to access the S3 bucket
resource "aws_cloudfront_origin_access_identity" "origin_access_identity" {
  comment = "Access identity for CloudFront to access S3 bucket"
}

# Policy for the S3 bucket to allow access only to CloudFront
resource "aws_s3_bucket_policy" "skatmate_policy" {
  bucket = aws_s3_bucket.skatmate_web_bucket.id
  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Principal = {
          AWS = aws_cloudfront_origin_access_identity.origin_access_identity.iam_arn
        },
        Action   = "s3:GetObject",
        Resource = "arn:aws:s3:::${aws_s3_bucket.skatmate_web_bucket.id}/*"
      }
    ]
  })
}

output "cloudfront_url" {
  value = aws_cloudfront_distribution.cdn.domain_name
}

output "api_endpoint" {
  value = "${aws_apigatewayv2_api.api_gateway.api_endpoint}/data"
}
