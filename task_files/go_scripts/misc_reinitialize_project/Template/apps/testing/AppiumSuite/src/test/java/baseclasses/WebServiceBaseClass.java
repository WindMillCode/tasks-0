package baseclasses;

import org.apache.hc.client5.http.classic.methods.HttpGet;
import org.apache.hc.client5.http.impl.classic.CloseableHttpClient;
import org.apache.hc.client5.http.impl.classic.CloseableHttpResponse;
import org.apache.hc.client5.http.impl.classic.HttpClientBuilder;
import org.apache.hc.core5.http.HttpStatus;
import org.testng.SkipException;
import org.testng.annotations.AfterClass;
import org.testng.annotations.BeforeClass;
import org.testng.annotations.BeforeMethod;

import java.io.IOException;

public class WebServiceBaseClass {

    protected CloseableHttpClient client;
    protected CloseableHttpResponse response;

    @BeforeClass
    public void setup() throws IOException {
        System.out.println("Runs once per class");
        client = HttpClientBuilder.create().build();
        response = client.execute(new HttpGet("https://api.github.com/lfkgfdlkgjslgj"));

        int actualStatusCode = response.getCode();
        if (actualStatusCode != HttpStatus.SC_OK) {
            throw new SkipException("Basic criteria failed, " +
                    "was expecting code 200, but got: " + actualStatusCode);
        }
    }

    @BeforeMethod
    public void setupMethod() {
        System.out.println("Runs before each @Test");
    }

    @AfterClass(alwaysRun = true)
    public void cleanup() throws IOException {
        response.close();
        client.close();
    }
}