// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

package com.google.crypto.tink.jwt;

import static com.google.common.truth.Truth.assertThat;
import static java.time.temporal.ChronoUnit.MILLIS;
import static org.junit.Assert.assertThrows;

import java.time.Clock;
import java.time.Duration;
import java.time.Instant;
import java.time.ZoneOffset;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.junit.runners.JUnit4;

/** Unit tests for JwtValidator */
@RunWith(JUnit4.class)
public final class JwtValidatorTest {

  @Test
  public void setNullValue_shouldThrow() throws Exception {
    assertThrows(NullPointerException.class, () -> JwtValidator.newBuilder().expectIssuer(null));
    assertThrows(NullPointerException.class, () -> JwtValidator.newBuilder().expectSubject(null));
    assertThrows(NullPointerException.class, () -> JwtValidator.newBuilder().expectAudience(null));
    assertThrows(NullPointerException.class, () -> JwtValidator.newBuilder().setClock(null));
    assertThrows(NullPointerException.class, () -> JwtValidator.newBuilder().setClockSkew(null));
  }

  @Test
  public void validate_expired_shouldThrow() throws Exception {
    Clock clock1 = Clock.systemUTC();
    // This token expires in 1 minute in the future.
    RawJwt token =
        RawJwt.newBuilder()
            .setExpiration(clock1.instant().plus(Duration.ofMinutes(1)))
            .build();

    // Move the clock to 2 minutes in the future.
    Clock clock2 = Clock.offset(clock1, Duration.ofMinutes(2));
    JwtValidator validator = JwtValidator.newBuilder().setClock(clock2).build();

    assertThrows(JwtInvalidException.class, () -> validator.validate(token));
  }

  @Test
  public void validate_notExpired_success() throws Exception {
    Clock clock = Clock.systemUTC();
    // This token expires in 1 minute in the future.
    Instant expiration = clock.instant().plus(Duration.ofMinutes(1));
    RawJwt unverified =
        RawJwt.newBuilder().setExpiration(expiration).build();
    JwtValidator validator = JwtValidator.newBuilder().build();
    VerifiedJwt token = validator.validate(unverified);

    assertThat(token.getExpiration()).isEqualTo(expiration.truncatedTo(MILLIS));
  }

  @Test
  public void validate_notExpired_clockSkew_success() throws Exception {
    Clock clock1 = Clock.systemUTC();
    // This token expires in 1 minutes in the future.
    Instant expiration = clock1.instant().plus(Duration.ofMinutes(1));
    RawJwt unverified =
        RawJwt.newBuilder().setExpiration(expiration).build();
    // A clock skew of 1 minute is allowed.
    JwtValidator validator = JwtValidator.newBuilder().setClockSkew(Duration.ofMinutes(1)).build();
    VerifiedJwt token = validator.validate(unverified);

    assertThat(token.getExpiration()).isEqualTo(expiration.truncatedTo(MILLIS));
  }

  @Test
  public void validate_tokenThatExpiresNow_shouldThrow() throws Exception {
    Instant expiration = Instant.ofEpochSecond(1234);
    Clock clock = Clock.fixed(expiration, ZoneOffset.UTC);
    RawJwt rawJwt =
        RawJwt.newBuilder().setExpiration(expiration).build();
    JwtValidator validator = JwtValidator.newBuilder().setClock(clock).build();

    assertThrows(JwtInvalidException.class, () -> validator.validate(rawJwt));
  }

  @Test
  public void validate_before_shouldThrow() throws Exception {
    Clock clock = Clock.systemUTC();
    // This token cannot be used until 1 minute in the future.
    Instant notBefore = clock.instant().plus(Duration.ofMinutes(1));
    RawJwt token =
        RawJwt.newBuilder().setNotBefore(notBefore).build();
    JwtValidator validator = JwtValidator.newBuilder().build();

    assertThrows(JwtInvalidException.class, () -> validator.validate(token));
  }

  @Test
  public void validate_notBefore_success() throws Exception {
    Clock clock1 = Clock.systemUTC();
    // This token cannot be used until 1 minute in the future.
    Instant notBefore = clock1.instant().plus(Duration.ofMinutes(1));
    RawJwt unverified =
        RawJwt.newBuilder().setNotBefore(notBefore).build();
    // Move the clock to 2 minutes in the future.
    Clock clock2 = Clock.offset(clock1, Duration.ofMinutes(2));
    JwtValidator validator = JwtValidator.newBuilder().setClock(clock2).build();
    VerifiedJwt token = validator.validate(unverified);

    assertThat(token.getNotBefore()).isEqualTo(notBefore.truncatedTo(MILLIS));
  }

  @Test
  public void validate_tokenWithNotBeforeIsNow_success() throws Exception {
    Instant notBefore = Instant.ofEpochSecond(1234);
    Clock clock = Clock.fixed(notBefore, ZoneOffset.UTC);
    RawJwt rawJwt =
        RawJwt.newBuilder().setNotBefore(notBefore).build();
    JwtValidator validator = JwtValidator.newBuilder().setClock(clock).build();
    VerifiedJwt token = validator.validate(rawJwt);
    assertThat(token.getNotBefore()).isEqualTo(notBefore);
  }

  @Test
  public void validate_notBefore_clockSkew_success() throws Exception {
    Clock clock1 = Clock.systemUTC();
    // This token cannot be used until 1 minute in the future.
    Instant notBefore = clock1.instant().plus(Duration.ofMinutes(1));
    RawJwt unverified =
        RawJwt.newBuilder().setNotBefore(notBefore).build();
    // A clock skew of 1 minute is allowed.
    JwtValidator validator = JwtValidator.newBuilder().setClockSkew(Duration.ofMinutes(1)).build();
    VerifiedJwt token = validator.validate(unverified);

    assertThat(token.getNotBefore()).isEqualTo(notBefore.truncatedTo(MILLIS));
  }

  @Test
  public void requireTypeHeaderButNoTypeHeaderInToken_shouldThrow() throws Exception {
    RawJwt token = RawJwt.newBuilder().build();
    JwtValidator validator = JwtValidator.newBuilder().expectTypeHeader("jwt").build();

    assertThrows(JwtInvalidException.class, () -> validator.validate(token));
  }

  @Test
  public void wrongTypeHeaderInToken_shouldThrow() throws Exception {
    RawJwt token =
        RawJwt.newBuilder().setTypeHeader("blah").build();
    JwtValidator validator = JwtValidator.newBuilder().expectTypeHeader("jwt").build();

    assertThrows(JwtInvalidException.class, () -> validator.validate(token));
  }

  @Test
  public void correctTypeHeaderInToken_success() throws Exception {
    RawJwt unverified =
        RawJwt.newBuilder().setTypeHeader("jwt").build();
    JwtValidator validator = JwtValidator.newBuilder().expectTypeHeader("jwt").build();
    VerifiedJwt token = validator.validate(unverified);
    assertThat(token.getTypeHeader()).isEqualTo("jwt");
  }

  @Test
  public void noTypeHeader_success() throws Exception {
    JwtValidator validator = JwtValidator.newBuilder().build();

    RawJwt tokenWithoutTypeHeader = RawJwt.newBuilder().build();
    validator.validate(tokenWithoutTypeHeader);
  }

  @Test
  public void typeHeaderInTokenButNoTypeHeaderSetInValidator_shouldThrow() throws Exception {
    JwtValidator validator = JwtValidator.newBuilder().build();

    RawJwt tokenWithTypeHeader = RawJwt.newBuilder().setTypeHeader("headerType").build();
    assertThrows(JwtInvalidException.class, () -> validator.validate(tokenWithTypeHeader));
  }

  @Test
  public void ignoreTypeHeaderSkipsValidationOfTypeHeader() throws Exception {
    JwtValidator validator = JwtValidator.newBuilder().ignoreTypeHeader().build();

    RawJwt tokenWithTypeHeader = RawJwt.newBuilder().setTypeHeader("headerType").build();
    validator.validate(tokenWithTypeHeader);
    RawJwt tokenWithoutTypeHeader = RawJwt.newBuilder().build();
    validator.validate(tokenWithoutTypeHeader);
  }

  @Test
  public void requireIssuerButNoIssuerInToken_shouldThrow() throws Exception {
    RawJwt token = RawJwt.newBuilder().build();
    JwtValidator validator = JwtValidator.newBuilder().expectIssuer("123").build();

    assertThrows(JwtInvalidException.class, () -> validator.validate(token));
  }

  @Test
  public void wrongIssuerInToken_shouldThrow() throws Exception {
    RawJwt token =
        RawJwt.newBuilder().setIssuer("blah").build();
    JwtValidator validator = JwtValidator.newBuilder().expectIssuer("123").build();

    assertThrows(JwtInvalidException.class, () -> validator.validate(token));
  }

  @Test
  public void correctIssuerInToken_success() throws Exception {
    RawJwt unverified =
        RawJwt.newBuilder().setIssuer("123").build();
    JwtValidator validator = JwtValidator.newBuilder().expectIssuer("123").build();
    VerifiedJwt token = validator.validate(unverified);
    assertThat(token.getIssuer()).isEqualTo("123");
  }

  @Test
  public void noIssuer_success() throws Exception {
    JwtValidator validator = JwtValidator.newBuilder().build();

    RawJwt tokenWithoutIssuer = RawJwt.newBuilder().build();
    validator.validate(tokenWithoutIssuer);
  }

  @Test
  public void issuerInTokenButNoIssuerSetInValidator_shouldThrow() throws Exception {
    JwtValidator validator = JwtValidator.newBuilder().build();

    RawJwt tokenWithIssuer = RawJwt.newBuilder().setIssuer("issuer").build();
    assertThrows(JwtInvalidException.class, () -> validator.validate(tokenWithIssuer));
  }

  @Test
  public void ignoreIssuerSkipsValidationOfIssuer() throws Exception {
    JwtValidator validator = JwtValidator.newBuilder().ignoreIssuer().build();

    RawJwt tokenWithIssuer = RawJwt.newBuilder().setIssuer("issuer").build();
    validator.validate(tokenWithIssuer);
    RawJwt tokenWithoutIssuer = RawJwt.newBuilder().build();
    validator.validate(tokenWithoutIssuer);
  }

  @Test
  public void requireSubjectButNoSubjectInToken_shouldThrow() throws Exception {
    RawJwt token = RawJwt.newBuilder().build();
    JwtValidator validator = JwtValidator.newBuilder().expectSubject("123").build();

    assertThrows(JwtInvalidException.class, () -> validator.validate(token));
  }

  @Test
  public void wrongSubjectInToken_shouldThrow() throws Exception {
    RawJwt token =
        RawJwt.newBuilder().setSubject("blah").build();
    JwtValidator validator = JwtValidator.newBuilder().expectSubject("123").build();

    assertThrows(JwtInvalidException.class, () -> validator.validate(token));
  }

  @Test
  public void correctSubjectInToken_success() throws Exception {
    RawJwt unverified =
        RawJwt.newBuilder().setSubject("123").build();
    JwtValidator validator = JwtValidator.newBuilder().expectSubject("123").build();
    VerifiedJwt token = validator.validate(unverified);

    assertThat(token.getSubject()).isEqualTo("123");
  }

  @Test
  public void noSubject_success() throws Exception {
    JwtValidator validator = JwtValidator.newBuilder().build();

    RawJwt tokenWithoutSubject = RawJwt.newBuilder().build();
    validator.validate(tokenWithoutSubject);
  }

  @Test
  public void subjectInTokenButNoSubjectSetInValidator_shouldThrow() throws Exception {
    JwtValidator validator = JwtValidator.newBuilder().build();

    RawJwt tokenWithSubject = RawJwt.newBuilder().setSubject("subject").build();
    assertThrows(JwtInvalidException.class, () -> validator.validate(tokenWithSubject));
  }

  @Test
  public void ignoreSubjectSkipsValidationOfSubject() throws Exception {
    JwtValidator validator = JwtValidator.newBuilder().ignoreSubject().build();

    RawJwt tokenWithSubject = RawJwt.newBuilder().setSubject("subject").build();
    validator.validate(tokenWithSubject);
    RawJwt tokenWithoutSubject = RawJwt.newBuilder().build();
    validator.validate(tokenWithoutSubject);
  }

  @Test
  public void requireAudienceButNoAudienceInToken_shouldThrow() throws Exception {
    RawJwt unverified = RawJwt.newBuilder().build();
    JwtValidator validator = JwtValidator.newBuilder().expectAudience("foo").build();

    assertThrows(JwtInvalidException.class, () -> validator.validate(unverified));
  }

  @Test
  public void audienceInTokenButNoAudienceSetInValidator_shouldThrow() throws Exception {
    RawJwt unverified =
        RawJwt.newBuilder().addAudience("foo").build();
    JwtValidator validator = JwtValidator.newBuilder().build();

    assertThrows(JwtInvalidException.class, () -> validator.validate(unverified));
  }

  @Test
  public void noAudience_success() throws Exception {
    RawJwt token = RawJwt.newBuilder().build();
    JwtValidator validator = JwtValidator.newBuilder().build();
    validator.validate(token);
  }

  @Test
  public void wrongAudienceInToken_shouldThrow() throws Exception {
    RawJwt unverified =
        RawJwt.newBuilder().addAudience("foo").build();
    JwtValidator validator = JwtValidator.newBuilder().expectAudience("bar").build();

    assertThrows(JwtInvalidException.class, () -> validator.validate(unverified));
  }

  @Test
  public void correctAudienceInToken_success() throws Exception {
    RawJwt unverified =
        RawJwt.newBuilder().addAudience("foo").build();
    JwtValidator validator = JwtValidator.newBuilder().expectAudience("foo").build();
    VerifiedJwt token = validator.validate(unverified);

    assertThat(token.getAudiences()).containsExactly("foo");
  }

  @Test
  public void correctAudienceInToken2_success() throws Exception {
    RawJwt unverified =
        RawJwt.newBuilder()
            .addAudience("foo")
            .addAudience("bar")
            .build();
    JwtValidator validator = JwtValidator.newBuilder().expectAudience("bar").build();
    VerifiedJwt token = validator.validate(unverified);

    assertThat(token.getAudiences()).containsExactly("foo", "bar");
  }

  @Test
  public void ignoreAudiencesSkipsValidationOfAudiences() throws Exception {
    JwtValidator validator = JwtValidator.newBuilder().ignoreAudiences().build();

    RawJwt tokenWithAudiences =
        RawJwt.newBuilder().addAudience("audience1").addAudience("audience2").build();
    validator.validate(tokenWithAudiences);
    RawJwt tokenWithoutAudience = RawJwt.newBuilder().build();
    validator.validate(tokenWithoutAudience);
  }

  @Test
  public void invalidValidators_fail() throws Exception {
    assertThrows(
        IllegalArgumentException.class,
        () -> JwtValidator.newBuilder().expectTypeHeader("a").ignoreTypeHeader().build());
    assertThrows(
        IllegalArgumentException.class,
        () -> JwtValidator.newBuilder().expectIssuer("a").ignoreIssuer().build());
    assertThrows(
        IllegalArgumentException.class,
        () -> JwtValidator.newBuilder().expectSubject("a").ignoreSubject().build());
    assertThrows(
        IllegalArgumentException.class,
        () -> JwtValidator.newBuilder().expectAudience("a").ignoreAudiences().build());
  }
}
