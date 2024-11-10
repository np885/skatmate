provider "aws" {
  region = "eu-central-1"
}

# Route 53: Hosted zone for the domain
resource "aws_route53_zone" "skatmate_zone" {
  name = "skatmate.de"
}

# SSL Certificate for CloudFront (ACM)
resource "aws_acm_certificate" "skatmate_cert" {
  domain_name       = "skatmate.de"
  validation_method = "DNS"

  # Add additional SAN for the www subdomain if you plan to support it
  subject_alternative_names = ["www.skatmate.de"]

  lifecycle {
    create_before_destroy = true
  }
}

# DNS validation record for the ACM certificate
resource "aws_route53_record" "cert_validation" {
  for_each = { for dvo in aws_acm_certificate.skatmate_cert.domain_validation_options : dvo.domain_name => dvo }
  zone_id  = aws_route53_zone.skatmate_zone.zone_id
  name     = each.value.resource_record_name
  type     = each.value.resource_record_type
  records  = [each.value.resource_record_value]
  ttl      = 60
}

# Certificate validation
resource "aws_acm_certificate_validation" "skatmate_cert_validation" {
  certificate_arn         = aws_acm_certificate.skatmate_cert.arn
  validation_record_fqdns = [for record in aws_route53_record.cert_validation : record.fqdn]
}

# S3 bucket for website content
resource "aws_s3_bucket" "skatmate_web_bucket" {
  bucket = "skatmate-web-bucket"
}

resource "aws_s3_bucket_website_configuration" "website_config" {
  bucket = aws_s3_bucket.skatmate_web_bucket.id

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "error.html"
  }
}

# CloudFront distribution for S3 bucket with custom domain and HTTPS
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

  aliases = ["skatmate.de", "www.skatmate.de"]

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
    acm_certificate_arn            = aws_acm_certificate.skatmate_cert.arn
    ssl_support_method              = "sni-only"
    minimum_protocol_version        = "TLSv1.2_2021"
  }
}

# CloudFront Origin Access Identity for S3 bucket access
resource "aws_cloudfront_origin_access_identity" "origin_access_identity" {
  comment = "Access identity for CloudFront to access S3 bucket"
}

# Bucket policy to allow CloudFront access to the S3 bucket
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

# Route 53 records for the domain
resource "aws_route53_record" "skatmate_root" {
  zone_id = aws_route53_zone.skatmate_zone.zone_id
  name    = "skatmate.de"
  type    = "A"

  alias {
    name                   = aws_cloudfront_distribution.cdn.domain_name
    zone_id                = aws_cloudfront_distribution.cdn.hosted_zone_id
    evaluate_target_health = false
  }
}

resource "aws_route53_record" "skatmate_www" {
  zone_id = aws_route53_zone.skatmate_zone.zone_id
  name    = "www.skatmate.de"
  type    = "A"

  alias {
    name                   = aws_cloudfront_distribution.cdn.domain_name
    zone_id                = aws_cloudfront_distribution.cdn.hosted_zone_id
    evaluate_target_health = false
  }
}

output "cloudfront_url" {
  value = aws_cloudfront_distribution.cdn.domain_name
}
