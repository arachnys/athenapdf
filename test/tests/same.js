const comparePdf = require("compare-pdf");
const chai = require("chai");
const expect = chai.expect;
const config = require("../config/config");

describe("Compare same Pdf Tests in Mocha + Chai", () => {
    it("Compare same pdfs by image", async () => {
        let comparisonResults = await new comparePdf(config)
            .actualPdfFile("new.pdf")
            .baselinePdfFile("baseline.pdf")
            .compare("byImage");
        expect(comparisonResults.status).to.equal("passed");
    });

    it("Compare pdfs by base64", async () => {
        let comparisonResults = await new comparePdf(config)
            .actualPdfFile("new.pdf")
            .baselinePdfFile("baseline.pdf")
            .compare("byBase64");
        expect(comparisonResults.status).to.equal("passed");
    });
});